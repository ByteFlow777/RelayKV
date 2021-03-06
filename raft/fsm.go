package raft

func (r *Raft) notifyCommit() {
	r.commitCh <- struct{}{}
}

type ApplyMsg struct {
	Data []byte
}

func (r *Raft) runFSM() {
	for {
		select {
		case <-r.commitCh:
			applied, committed := r.getLastApplied()+1, r.getCommitIndex()
			for ; applied <= committed; applied++ {
				entry, err := r.logStore.GetEntry(applied)
				if err != nil {
					panic(err)
				}
				r.logger.Infof("%v apply: %v", r.me(), entry)
				r.applyCh <- &ApplyMsg{
					entry.Data,
				}
			}
			r.setLastApplied(committed)
		}
	}
}
