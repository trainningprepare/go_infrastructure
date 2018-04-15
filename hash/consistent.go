package consistent

import (
	"sort"
	"fmt"
	"hash/fnv"
)

type Bucket struct{
	Name               string
	TargetRingPosition uint64
}

type ConsitentRing struct{
	Range   uint64   //scope 范围 模数
	Buckets []*Bucket  //环上的控制节点
}

func (this *ConsitentRing) Len() int {
	return len(this.Buckets)
}

func (this *ConsitentRing) Swap(i,j int)  {
     this.Buckets[i],this.Buckets[j] = this.Buckets[j],this.Buckets[i]
}

func (this *ConsitentRing) Less(i,j int) bool {
	return (this.Buckets[i].TargetRingPosition) < (this.Buckets[j].TargetRingPosition)
}

func (this *ConsitentRing) AddBucket(name string) {
	hashV:=fnv.New64()
	hashV.Write([]byte(name))
	this.Buckets = append(this.Buckets,&Bucket{name,hashV.Sum64()%this.Range})
}

func (this *ConsitentRing) DelNode(node uint64)  {
	return
}

func (this *ConsitentRing)DumpNodesRange() *ConsitentRing  {
	sort.Sort(this)
	var value *Bucket
	for i := 0; i <int(this.Range); i++ {
		value = this.Buckets[i]
		fmt.Printf("bucketName:%s ,                 bucket.position:%v \n",value.Name,value.TargetRingPosition)
	}

	
	return this
}

func (this *ConsitentRing)FindBucketByKey(key string) *Bucket {
    hasv:= fnv.New64()
    hasv.Write([]byte(key))
    keyPos:=hasv.Sum64()%this.Range
    start:= (sort.Search(len(this.Buckets), func(i int) bool {
		return this.Buckets[i].TargetRingPosition >=keyPos
	}) - 1) + len(this.Buckets)
    start = start % len(this.Buckets)
    fmt.Println(start)
    return  this.Buckets[start]
}

func CreateConsitentRing(blockCount uint64) *ConsitentRing {
    return &ConsitentRing{blockCount,nil}
}