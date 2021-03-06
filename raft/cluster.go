package raft

import (
	"fmt"
)

type ServerID uint64

type ServerAddress string

const (
	NoneID      ServerID      = 0
	NoneAddress ServerAddress = ""
)

type Server struct {
	ServerID
	ServerAddress
}

func (s Server) String() string {
	return fmt.Sprintf("Server{ ID: %v, Address: %v }",
		s.ServerID, s.ServerAddress)
}

type Cluster []Server

func NewCluster(address []string) Cluster {
	var clusters Cluster
	for id, add := range address {
		clusters = append(clusters, Server{
			ServerID:      ServerID(id),
			ServerAddress: ServerAddress(add),
		})
	}
	return clusters
}

func (c *Cluster) size() uint64 {
	return uint64(len(*c))
}

func (c *Cluster) quorum() uint64 {
	return uint64(len(*c)) / 2
}

func (c *Cluster) at(index uint64) Server {
	if index > c.size() {
		panic("index out of bound of the cluster")
	}
	return (*c)[index]
}

func (c *Cluster) visit(f func(s Server), async bool) {
	for _, server := range *c {
		if async {
			go f(server)
		} else {
			f(server)
		}
	}
}
