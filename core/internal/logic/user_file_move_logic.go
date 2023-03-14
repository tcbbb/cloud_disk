package logic

import (
	"cloud-disk/core/models"
	"context"
	"errors"

	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFileMoveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileMoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileMoveLogic {
	return &UserFileMoveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// 首先查看要移动到的目标文件夹，ParentIdnetity是否存在。若存在，将Idnetity对应的目标文件中的parentId转为目标文件夹的ID
func (l *UserFileMoveLogic) UserFileMove(req *types.UserFileMoveRequest, userIdentity string) (resp *types.UserFileMoveReply, err error) {
	//parentID
	parentData := new(models.UserRepository)
	// 获取目标文件夹信息
	has, err := l.svcCtx.Engine.Where("identity = ? AND user_identity = ?", req.ParentIdnetity, userIdentity).Get(parentData)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("文件夹不存在")
	}
	// 更新记录的 ParentID
	_, err = l.svcCtx.Engine.Where("identity = ?", req.Idnetity).Update(models.UserRepository{
		ParentId: int64(parentData.Id),
	})
	return
}
