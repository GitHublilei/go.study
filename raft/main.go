package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// 实现3节点选举

// 改造代码成分布式选举代码， 加入rpc调用

// 演示完整代码 自动选主 日志复制

// 定义3节点常量

const raftCount = 3

// Leader leader对象
type Leader struct {
	Term     int // 任期
	LeaderID int // 编号
}

// Raft ...
type Raft struct {
	mu              sync.Mutex // 锁
	me              int        //节点编号
	currentTerm     int        // 当前任期
	votedFor        int        // 为哪个节点投票
	state           int        // 0 follower 1 candidate 2 leader
	lastMessageTime int64      // 发送最后一条数据的时间
	currentLeader   int        // 设置当前节点的领导
	message         chan bool  // 节点间发信息的通道
	electCh         chan bool  // 选举通道
	heartBeat       chan bool  // 心跳信号的通道
	heartbreatRe    chan bool  // 返回心跳的通道
	timeout         int        // 超时时间
}

var leader = Leader{0, -1}

// Mack 创建节点
func Mack(me int) *Raft {
	rf := &Raft{}
	rf.me = me
	// -1代表谁都不投， 此时节点刚创建
	rf.votedFor = -1
	rf.state = 0
	rf.timeout = 0
	rf.currentLeader = -1
	// 节点任期
	rf.setTerm(0)
	// 初始化通道
	rf.message = make(chan bool)
	rf.electCh = make(chan bool)
	rf.heartBeat = make(chan bool)
	rf.heartbreatRe = make(chan bool)
	// 设置随机种子
	rand.Seed(time.Now().UnixNano())
	// 选举的协程
	go rf.election()
	// 心跳检测的协程
	go rf.sendLeaderHeartBeat()
	return rf
}

func (r *Raft) setTerm(term int) {
	r.currentTerm = term
}

func (r *Raft) election() {
	// 设置标记， 判断是否选出了leader
	var result bool
	for {
		// 设置超时， 150 到300之间
		timeout := randRange(150, 300)
		r.lastMessageTime = millisecond()
		select {
		// 延迟等待1毫秒
		case <-time.After(time.Duration(timeout) * time.Millisecond):
			fmt.Println("当前节点状态：", r.state)
		}
		result = false
		for !result {
			// 选主逻辑
		}
	}
}

// 实现选主的逻辑
func (r *Raft) electionOneRound(leader *Leader) bool {
	// 定义超时
	var timeout int64
	timeout = 100
	//投票数量
	var vote int
	// 定义是否开始心跳信号的产生
	var triggerHearbeat bool
	// 时间
	last := millisecond()
	// 用于返回指
	success := false

	// 给当前节点变成cadidate
	r.mu.Lock()
	// 修改状态
	r.becomeCandidate()
	r.mu.Unlock()
	fmt.Println("start electing leader")
	for {
		// 遍历所有节点拉选票
		for i := 0; i < raftCount; i++ {
			if i != r.me {
				go func() {
					if leader.LeaderID < 0 {
						r.electCh <- true
					}
				}()
			}
		}
		// 设置投票数量
		vote = 1
		// 遍历
		for i := 0; i < raftCount; i++ {
			// 计算投票数量
			select {
			case ok := <-r.electCh:
				if ok {
					vote++
					success = vote > raftCount/2
					if success && !triggerHearbeat {
						//变化成主节点，选举成功了
						// 开始触发心跳信号检测
						triggerHearbeat = true
						r.mu.Lock()
						r.becomeLeader()
						r.mu.Unlock()
						// 由leader向其它节点发送心跳信号
						fmt.Println(r.me, "节号点成为了心跳信号了")
					}
				}
			}
		}

		// 做最后校验工作
		// 若不超时，且票数大于一半，则选举成功， break
		if timeout+last < millisecond() || (vote > raftCount/2 || r.currentLeader > -1) {
			break
		} else {
			// 等待操作
			select {
			case <-time.After(time.Duration(10) * time.Millisecond):
			}
		}
	}
	return success
}

// 修改状态candidate
func (r *Raft) becomeCandidate() {
	r.state = 1
	r.setTerm(r.currentTerm + 1)
	r.votedFor = r.me
	r.currentLeader = -1
}

// 变成leader
func (r *Raft) becomeLeader() {
	r.state = 2
	r.currentLeader = r.me
}

// leader节点发送心跳信号
func (r *Raft) sendLeaderHeartBeat() {
}

// 随机值
func randRange(min, max int64) int64 {
	return rand.Int63n(max-min) + min
}

// 获取当前时间， 发送最后一条数据的时间
func millisecond() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func main() {
	// 有3个节点，最初都是follower
	// 若有candidate状态， 进行投票拉票
	// 会产生leader
	// 创建3个节点
	for i := 0; i < raftCount; i++ {
		Mack(i)
	}

	for {
	}
}
