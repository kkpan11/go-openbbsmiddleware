package apitypes

import "github.com/Ptt-official-app/go-pttbbs/ptttype"

type UserBoardInfo struct {
	Read bool
	Fav  bool
	Stat ptttype.BoardStatAttr
}
