package repository

import (
	"time"

	"github.com/8treenet/freedom/example/fshop/application/entity"
	"github.com/8treenet/freedom/example/fshop/application/object"

	"github.com/8treenet/freedom"
)

func init() {
	freedom.Booting(func(initiator freedom.Initiator) {
		initiator.BindRepository(func() *Goods {
			return &Goods{}
		})
	})
}

var _ GoodsRepo = new(Goods)

// Goods .
type Goods struct {
	freedom.Repository
}

func (repo *Goods) Find(id int) (goodsEntity *entity.Goods, e error) {
	goodsEntity = &entity.Goods{}
	e = findGoodsByPrimary(repo, goodsEntity, id)
	if e != nil {
		return
	}
	return
}

func (repo *Goods) Finds(ids []int) (entitys []*entity.Goods, e error) {
	var primarys []interface{}
	for i := 0; i < len(ids); i++ {
		primarys = append(primarys, ids[i])
	}
	e = findGoodssByPrimarys(repo, &entitys, primarys...)
	return
}

func (repo *Goods) FindsByPage(page, pageSize int, tag string) (entitys []*entity.Goods, e error) {
	build := repo.NewORMDescBuilder("id").NewPageBuilder(page, pageSize)
	e = findGoodss(repo, object.Goods{Tag: tag}, &entitys, build)
	// 分页器读取总数 fmt.Println(build.TotalPage())
	return
}

func (repo *Goods) Save(entity *entity.Goods) error {
	_, e := saveGoods(repo, &entity.Goods)
	return e
}

func (repo *Goods) New(name, tag string, price, stock int) (entityGoods *entity.Goods, e error) {
	goods := object.Goods{Name: name, Price: price, Stock: stock, Tag: tag, Created: time.Now(), Updated: time.Now()}

	_, e = createGoods(repo, &goods)
	if e != nil {
		return
	}
	entityGoods = &entity.Goods{Goods: goods}
	repo.InjectEntity(entityGoods)
	return
}
