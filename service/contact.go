package service

import (
	"errors"
	"fmt"
	"time"

	"../model"
)

type ContactService struct {
}

//自动添加好友
func (service *ContactService) AddFriend(fromId, toId int64) error {
	//如果加自己
	if fromId == toId {
		return errors.New("不能添加自己为好友啊")
	}
	//判断是否已经加了好友
	tmp := model.Contact{}
	//查询是否已经是好友
	// 条件的链式操作
	DbEngine.Where("fromId = ?", fromId).
		And("toId = ?", toId).
		And("cate = ?", model.CONCAT_CATE_USER).
		Get(&tmp)
	//获得1条记录
	//count()
	//如果存在记录说明已经是好友了不加
	if tmp.Id > 0 {
		return errors.New("该用户已经被添加过啦")
	}

	//事务,
	session := DbEngine.NewSession()
	session.Begin()
	//插自己的
	_, e2 := session.InsertOne(model.Contact{
		FromId:   fromId,
		ToId:     toId,
		Cate:     model.CONCAT_CATE_USER,
		Createat: time.Now(),
	})
	//插对方的
	_, e3 := session.InsertOne(model.Contact{
		FromId:   toId,
		ToId:     fromId,
		Cate:     model.CONCAT_CATE_USER,
		Createat: time.Now(),
	})
	//没有错误
	if e2 == nil && e3 == nil {
		//提交
		session.Commit()
		return nil
	} else {
		//回滚,说明e2或e3至少有个err
		session.Rollback()
		if e2 != nil {
			return e2
		} else {
			return e3
		}
	}
}

//查找好友
func (service *ContactService) SearchFriend(userId int64) []model.User {
	conconts := make([]model.Contact, 0)
	objIds := make([]int64, 0)
	DbEngine.Where("from_id = ? and cate = ?", userId, model.CONCAT_CATE_USER).Find(&conconts)
	for _, v := range conconts {
		objIds = append(objIds, v.ToId)
	}
	fmt.Println(userId, conconts, objIds)
	coms := make([]model.User, 0)
	if len(objIds) == 0 {
		return coms
	}
	DbEngine.In("id", objIds).Find(&coms)
	return coms
}

// 用户加入的群
func (service *ContactService) SearchComunity(userId int64) []model.Community {
	conconts := make([]model.Contact, 0)
	comIds := make([]int64, 0)

	DbEngine.Where("ownerid = ? and cate = ?", userId, model.CONCAT_CATE_COMUNITY).Find(&conconts)
	for _, v := range conconts {
		comIds = append(comIds, v.ToId)
	}
	coms := make([]model.Community, 0)
	if len(comIds) == 0 {
		return coms
	}
	DbEngine.In("id", comIds).Find(&coms)
	return coms
}

//加群
func (service *ContactService) JoinCommunity(userId, comId int64) error {
	cot := model.Contact{
		FromId: userId,
		ToId:   comId,
		Cate:   model.CONCAT_CATE_COMUNITY,
	}
	DbEngine.Get(&cot)
	if cot.Id == 0 {
		cot.Createat = time.Now()
		_, err := DbEngine.InsertOne(cot)
		return err
	} else {
		return nil
	}

}

func (service *ContactService) SearchComunityIds(userId int64) (comIds []int64) {
	//todo 获取用户全部群ID
	conconts := make([]model.Contact, 0)
	comIds = make([]int64, 0)

	DbEngine.Where("ownerid = ? and cate = ?", userId, model.CONCAT_CATE_COMUNITY).Find(&conconts)
	for _, v := range conconts {
		comIds = append(comIds, v.ToId)
	}
	return comIds
}
