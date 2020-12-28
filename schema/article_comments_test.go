package schema

import (
	"testing"

	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/Ptt-official-app/go-pttbbs/testutil"
	"go.mongodb.org/mongo-driver/bson"
)

func TestUpdateArticleComments(t *testing.T) {
	setupTest()
	defer teardownTest()

	defer Article_c.Drop()

	articleContent := &ArticleContentInfo{
		ContentMD5:          "test1",
		IP:                  "127.0.0.1",
		Host:                "localhost",
		BBS:                 "ptt",
		ContentUpdateNanoTS: types.NanoTS(1234567890000000000),
	}

	_ = UpdateArticleContentInfo(bbs.BBoardID("board0"), bbs.ArticleID("article0"), articleContent)

	articleComments0 := &ArticleComments{
		BBoardID:             bbs.BBoardID("board0"),
		ArticleID:            bbs.ArticleID("article0"),
		CommentsUpdateNanoTS: types.NanoTS(1234567890000000000),
	}

	articleComments1 := &ArticleComments{
		BBoardID:             bbs.BBoardID("board1"),
		ArticleID:            bbs.ArticleID("article1"),
		CommentsUpdateNanoTS: types.NanoTS(1234567890000000000),
	}

	type args struct {
		articleComments *ArticleComments
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			args: args{articleComments: articleComments0},
		},
		{
			args:    args{articleComments: articleComments1},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := UpdateArticleComments(tt.args.articleComments)

			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateArticleComments() error = %v, wantErr %v", err, tt.wantErr)
			}

			if err != nil {
				return
			}

			article := &Article{}
			query := bson.M{
				ARTICLE_BBOARD_ID_b:  tt.args.articleComments.BBoardID,
				ARTICLE_ARTICLE_ID_b: tt.args.articleComments.ArticleID,
			}
			_ = Article_c.FindOne(query, article, nil)

			testutil.TDeepEqual(t, "CommentsUpdateNanoTS", article.CommentsUpdateNanoTS, tt.args.articleComments.CommentsUpdateNanoTS)
		})
	}
}