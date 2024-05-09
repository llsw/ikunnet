package balance

import (
	"context"
	"fmt"

	"github.com/cloudwego/kitex/pkg/discovery"
	"github.com/cloudwego/kitex/pkg/loadbalance"
	transport "github.com/llsw/ikunet/internal/kitex_gen/transport"
	cmap "github.com/orcaman/concurrent-map"
)

var (
	_         loadbalance.Loadbalancer = &Balancer{}
	pic       loadbalance.Picker       = &picker{}
	cache     cmap.ConcurrentMap       = cmap.New()
	whiteList cmap.ConcurrentMap       = cmap.New()
)

const (
	TAG_CLUSTER  = "cluster"
	TAG_VERSION  = "version"
	TAG_ID       = "id"
	TAG_TYPE     = "type"
	TAG_MAINTAIN = "maintain"
)

type picker struct {
	dr *discovery.Result
}

func getNewVer(uuid string, dr *discovery.Result) string {
	var isWl bool
	if mv, ok := whiteList.Get(dr.CacheKey); ok {
		wl := mv.(map[string]bool)
		if _, ok := wl[uuid]; ok {
			isWl = true
		}
	}
	var ver string
	if vt, ok := dr.Instances[0].Tag(TAG_VERSION); ok {
		ver = vt
	}

	for k := 1; k < len(dr.Instances); k++ {
		v := dr.Instances[k]
		if _, ok := v.Tag(TAG_MAINTAIN); ok {
			if !isWl {
				continue
			}
		}
		// TODO 做负载均衡
		// TODO 如果有TAG_TYPE, 还要判断是有状态还是无状态，有状态需要根据查缓存，找出TAG_ID 看是原来走的是哪个
		if vt, ok := v.Tag(TAG_VERSION); ok {
			if vt > ver {
				ver = vt
			}
		}
	}
	return ver
}

func (p *picker) Next(ctx context.Context, request interface{}) discovery.Instance {
	req := request.(*transport.Transport)
	cmd := req.GetCmd()
	uuid := req.GetMeta().GetUuid()
	var ver string
	vck := fmt.Sprintf("%s_%s", uuid, p.dr.CacheKey)
	if cv, ok := cache.Get(vck); !ok {
		ver = getNewVer(uuid, p.dr)
	} else {
		ver = cv.(string)
	}

	if _, ok := p.dr.Instances[0].Tag(GetBlCallKey(cmd)); ok {
		newVer := getNewVer(uuid, p.dr)
		if ver != newVer {
			cache.Set(vck, newVer)
			ver = newVer
		}
	}

	for _, v := range p.dr.Instances {
		// TODO: 负载均衡
		if iv, ok := v.Tag(TAG_MAINTAIN); ok {
			if iv == ver {
				return v
			}
		}
	}

	return nil
}

type Balancer struct {
}

func (b *Balancer) GetPicker(dr discovery.Result) loadbalance.Picker {
	// 每次都要赋值，防止服务发现结果改变了
	pic.(*picker).dr = &dr
	return pic
}

func (b *Balancer) Name() string {
	return "ikunet_balancer"
}

func GetBlCallKey(call string) string {
	return fmt.Sprintf("blcall-%s", call)
}
