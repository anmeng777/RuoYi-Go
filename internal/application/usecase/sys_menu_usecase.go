// Copyright (c) [2024] K. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file. or see：https://github.com/Kun-GitHub/RuoYi-Go/blob/main/LICENSE
// Author: K. See：https://github.com/Kun-GitHub/RuoYi-Go or https://gitee.com/gitee_kun/RuoYi-Go
// Email: hot_kun@hotmail.com or 867917691@qq.com

package usecase

import (
	"RuoYi-Go/internal/common"
	"RuoYi-Go/internal/domain/model"
	"RuoYi-Go/internal/ports/input"
	"RuoYi-Go/internal/ports/output"
	"RuoYi-Go/pkg/cache"
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
)

type SysMenuService struct {
	repo   output.SysMenuRepository
	cache  *cache.FreeCacheClient
	logger *zap.Logger
}

func NewSysMenuService(repo output.SysMenuRepository, cache *cache.FreeCacheClient, logger *zap.Logger) input.SysMenuService {
	return &SysMenuService{repo: repo, cache: cache, logger: logger}
}

func (this *SysMenuService) QueryMenuByID(id int64) (*model.SysMenu, error) {
	structEntity := &model.SysMenu{}
	// 尝试从缓存中获取
	userBytes, err := this.cache.Get([]byte(fmt.Sprintf("MenuID:%d", id)))
	if err == nil {
		// 缓存命中
		err = json.Unmarshal(userBytes, &structEntity)
		if err == nil && structEntity.MenuID != 0 {
			// 缓存命中
			return structEntity, nil
		}
	}

	structEntity, err = this.repo.QueryMenuByID(id)
	if err != nil {
		this.logger.Error("QueryMenuByID", zap.Error(err))
		return nil, err
	} else if structEntity.MenuID != 0 {
		// 序列化用户对象并存入缓存
		userBytes, err = json.Marshal(structEntity)
		if err == nil {
			this.cache.Set([]byte(fmt.Sprintf("MenuID:%d", structEntity.MenuID)), userBytes, common.EXPIRESECONDS) // 第三个参数是过期时间，0表示永不过期
		}
		return structEntity, nil
	}

	this.logger.Debug("查询信息失败", zap.Error(err))
	return nil, fmt.Errorf("查询信息失败", zap.Error(err))
}

func (this *SysMenuService) QueryMenuList(post *model.SysMenuRequest) ([]*model.SysMenu, error) {
	structEntity := make([]*model.SysMenu, 0)

	structEntity, err := this.repo.QueryMenuList(post)
	if err != nil {
		this.logger.Error("QueryMenuList", zap.Error(err))
		return nil, err
	} else {
		return structEntity, nil
	}
}

func (this *SysMenuService) QueryMenuPage(pageReq common.PageRequest, r *model.SysMenuRequest) ([]*model.SysMenu, int64, error) {
	data, total, err := this.repo.QueryMenuPage(pageReq, r)
	if err != nil {
		this.logger.Error("QueryMenuPage", zap.Error(err))
		return nil, 0, err
	}

	return data, total, nil
}

func (this *SysMenuService) AddMenu(post *model.SysMenu) (*model.SysMenu, error) {
	data, err := this.repo.AddMenu(post)
	if err != nil {
		this.logger.Error("AddMenu", zap.Error(err))
		return nil, err
	}
	if data != nil && data.MenuID != 0 {
		// 序列化用户对象并存入缓存
		userBytes, err := json.Marshal(data)
		if err == nil {
			this.cache.Set([]byte(fmt.Sprintf("MenuID:%d", data.MenuID)), userBytes, common.EXPIRESECONDS) // 第三个参数是过期时间，0表示永不过期
		}
	}
	return data, nil
}

func (this *SysMenuService) EditMenu(post *model.SysMenu) (*model.SysMenu, int64, error) {
	data, result, err := this.repo.EditMenu(post)
	if err != nil {
		this.logger.Error("EditMenu", zap.Error(err))
		return nil, 0, err
	}
	if data != nil && data.MenuID != 0 && result == 1 {
		// 序列化用户对象并存入缓存
		userBytes, err := json.Marshal(data)
		if err == nil {
			this.cache.Set([]byte(fmt.Sprintf("MenuID:%d", data.MenuID)), userBytes, common.EXPIRESECONDS) // 第三个参数是过期时间，0表示永不过期
		}
	}
	return data, result, nil
}

func (this *SysMenuService) DeleteMenuById(id int64) (int64, error) {
	result, err := this.repo.DeleteMenuById(id)
	if err != nil {
		this.logger.Error("DeleteMenuById", zap.Error(err))
		return 0, err
	}
	if result > 0 {
		this.cache.Del(fmt.Sprintf("MenuID:%d", id))
	}
	return result, nil
}
