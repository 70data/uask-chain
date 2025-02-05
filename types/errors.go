package types

import "errors"

var (
	ErrQuestionNotFound = errors.New("question not found")
	ErrAnswerNotFound   = errors.New("answer not found")
	ErrCommentNotFound  = errors.New("comment not found")

	ErrNoPermission = errors.New("no permission")

	ErrRewardNotEnough = errors.New("reward not enough")
	ErrRewardIllegal   = errors.New("reward is illegal")

	ErrFileNotfound = errors.New("file-content not found")
)
