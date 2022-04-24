package request

import "snowy-video-serve/model"

type SysOperationRecordSearch struct {
	model.SysOperationRecord
	PageInfo
}
