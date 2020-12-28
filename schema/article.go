package schema

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/db"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/ptttype"
)

var (
	Article_c *db.Collection
)

type Article struct {
	BBoardID   bbs.BBoardID     `bson:"bid"`                 //
	ArticleID  bbs.ArticleID    `bson:"aid"`                 //
	IsDeleted  bool             `bson:"deleted,omitempty"`   //
	Filename   string           `bson:"filename"`            //
	CreateTime types.NanoTS     `bson:"create_time_nano_ts"` //
	MTime      types.NanoTS     `bson:"mtime_nano_ts"`       //
	Recommend  int              `bson:"recommend"`           //
	Owner      bbs.UUserID      `bson:"owner"`               //
	Title      string           `bson:"title"`               //
	Money      int              `bson:"money"`               //
	Class      string           `bson:"class"`               //
	Filemode   ptttype.FileMode `bson:"mode"`                //

	UpdateNanoTS types.NanoTS `bson:"update_nano_ts"` //used by article-summary

	ContentMTime        types.NanoTS    `bson:"content_mtime_nano_ts"` //
	ContentMD5          string          `bson:"content_md5"`
	Content             [][]*types.Rune `bson:"content"` //
	IP                  string          `bson:"ip"`      //
	Host                string          `bson:"host"`    //ip 的中文呈現, 外國則為國家.
	BBS                 string          `bson:"bbs"`     //
	ContentUpdateNanoTS types.NanoTS    `bson:"content_update_nano_ts"`

	FirstCommentsMD5          string       `bson:"first_comments_md5"`
	FirstCommentsLastTime     types.NanoTS `bson:"first_comments_last_time_nano_ts"`
	FirstCommentsUpdateNanoTS types.NanoTS `bson:"first_comments_update_nano_ts"`

	CommentsUpdateNanoTS types.NanoTS `bson:"comments_update_nano_ts"`
}

var (
	EMPTY_ARTICLE = &Article{}
)

var ( //bson-name
	ARTICLE_BBOARD_ID_b   = getBSONName(EMPTY_ARTICLE, "BBoardID")
	ARTICLE_ARTICLE_ID_b  = getBSONName(EMPTY_ARTICLE, "ArticleID")
	ARTICLE_IS_DELETED_b  = getBSONName(EMPTY_ARTICLE, "IsDeleted")
	ARTICLE_FILENAME_b    = getBSONName(EMPTY_ARTICLE, "Filename")
	ARTICLE_CREATE_TIME_b = getBSONName(EMPTY_ARTICLE, "CreateTime")
	ARTICLE_MTIME_b       = getBSONName(EMPTY_ARTICLE, "MTime")
	ARTICLE_RECOMMEND_b   = getBSONName(EMPTY_ARTICLE, "Recommend")
	ARTICLE_OWNER_b       = getBSONName(EMPTY_ARTICLE, "Owner")
	ARTICLE_TITLE_b       = getBSONName(EMPTY_ARTICLE, "Title")
	ARTICLE_MONEY_b       = getBSONName(EMPTY_ARTICLE, "Money")
	ARTICLE_CLASS_b       = getBSONName(EMPTY_ARTICLE, "Class")
	ARTICLE_FILEMODE_b    = getBSONName(EMPTY_ARTICLE, "Filemode")

	ARTICLE_UPDATE_NANO_TS_b = getBSONName(EMPTY_ARTICLE, "UpdateNanoTS")

	ARTICLE_CONTENT_MTIME_b = getBSONName(EMPTY_ARTICLE, "ContentMTime")
	ARTICLE_CONTENT_MD5_B   = getBSONName(EMPTY_ARTICLE, "ContentMD5")
	ARTICLE_CONTENT_b       = getBSONName(EMPTY_ARTICLE, "Content")
	ARTICLE_IP_b            = getBSONName(EMPTY_ARTICLE, "IP")
	ARTICLE_HOST_b          = getBSONName(EMPTY_ARTICLE, "Host")
	ARTICLE_BBS_b           = getBSONName(EMPTY_ARTICLE, "BBS")

	ARTICLE_CONTENT_UPDATE_NANO_TS_b = getBSONName(EMPTY_ARTICLE, "ContentUpdateNanoTS")

	ARTICLE_FIRST_COMMENTS_MD5_b       = getBSONName(EMPTY_ARTICLE, "FirstCommentsMD5")
	ARTICLE_FIRST_COMMENTS_LAST_TIME_b = getBSONName(EMPTY_ARTICLE, "FirstCommentsLastTime")

	ARTICLE_FIRST_COMMENTS_UPDATE_NANO_TS_b = getBSONName(EMPTY_ARTICLE, "FirstCommentsUpdateNanoTS")

	ARTICLE_COMMENTS_UPDATE_NANO_TS_b = getBSONName(EMPTY_ARTICLE, "CommentsUpdateNanoTS")
)

func assertArticleFields() error {
	if err := assertFields(EMPTY_ARTICLE, EMPTY_ARTICLE_QUERY); err != nil {
		return err
	}

	//article_content_info
	if err := assertFields(EMPTY_ARTICLE, EMPTY_ARTICLE_CONTENT_INFO); err != nil {
		return err
	}

	//article_content_mtime
	if err := assertFields(EMPTY_ARTICLE, EMPTY_ARTICLE_CONTENT_MTIME); err != nil {
		return err
	}

	//article_detail_summary
	if err := assertFields(EMPTY_ARTICLE, EMPTY_ARTICLE_DETAIL_SUMMARY); err != nil {
		return err
	}

	//article_first_comments
	if err := assertFields(EMPTY_ARTICLE, EMPTY_ARTICLE_FIRST_COMMENTS); err != nil {
		return err
	}

	//article_summary
	if err := assertFields(EMPTY_ARTICLE, EMPTY_ARTICLE_SUMMARY); err != nil {
		return err
	}

	return nil
}

type ArticleQuery struct {
	BBoardID  bbs.BBoardID  `bson:"bid"`
	ArticleID bbs.ArticleID `bson:"aid"`
	IsDeleted interface{}   `bson:"deleted,omitempty"` //
}

var (
	EMPTY_ARTICLE_QUERY = &ArticleQuery{}
)