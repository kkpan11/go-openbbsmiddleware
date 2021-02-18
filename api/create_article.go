package api

import (
	"github.com/Ptt-official-app/go-openbbsmiddleware/apitypes"
	"github.com/Ptt-official-app/go-openbbsmiddleware/dbcs"
	"github.com/Ptt-official-app/go-openbbsmiddleware/types"
	"github.com/Ptt-official-app/go-openbbsmiddleware/utils"
	pttbbsapi "github.com/Ptt-official-app/go-pttbbs/api"
	"github.com/Ptt-official-app/go-pttbbs/bbs"
	"github.com/gin-gonic/gin"
)

const CREATE_ARTICLE_R = "/board/:bid/article"

type CreateArticleParams struct {
	PostType string          `json:"class" form:"class" url:"class"`
	Title    string          `json:"title" form:"title" url:"title"`
	Content  [][]*types.Rune `json:"content" form:"content" url:"content"`
}

type CreateArticlePath struct {
	BoardID bbs.BBoardID `uri:"bid" binding:"required"`
}

type CreateArticleResult *apitypes.ArticleSummary

func CreateArticleWrapper(c *gin.Context) {
	params := &CreateArticleParams{}
	path := &CreateArticlePath{}
	LoginRequiredPathJSON(CreateArticle, params, path, c)
}

func CreateArticle(remoteAddr string, userID bbs.UUserID, params interface{}, path interface{}, c *gin.Context) (result interface{}, statusCode int, err error) {

	theParams, ok := params.(*CreateArticleParams)
	if !ok {
		return nil, 400, ErrInvalidParams
	}

	thePath, ok := path.(*CreateArticlePath)
	if !ok {
		return nil, 400, ErrInvalidPath
	}

	theType := types.Utf8ToBig5(theParams.PostType)
	theTitle := types.Utf8ToBig5(theParams.Title)
	content := dbcs.Utf8ToBig5(theParams.Content)

	//backend
	theParams_b := &pttbbsapi.CreateArticleParams{
		PostType: theType,
		Title:    theTitle,
		Content:  content,
	}
	var result_b pttbbsapi.CreateArticleResult

	urlMap := map[string]string{
		"bid": string(thePath.BoardID),
	}
	url := utils.MergeURL(urlMap, pttbbsapi.CREATE_ARTICLE_R)
	statusCode, err = utils.BackendPost(c, url, theParams_b, nil, &result_b)
	if err != nil || statusCode != 200 {
		return nil, statusCode, err
	}

	//update to db
	theList_b := []*bbs.ArticleSummary{(*bbs.ArticleSummary)(result_b)}
	updateNanoTS := types.NowNanoTS()
	articleSummaries_db, _, err := deserializeArticlesAndUpdateDB(userID, thePath.BoardID, theList_b, updateNanoTS)
	if err != nil {
		return nil, 500, err
	}

	articleSummary_db := articleSummaries_db[0]

	dummyContent := [][]*types.Rune{}

	setUserReadArticle(dummyContent, userID, articleSummary_db.ArticleID, updateNanoTS)

	articleSummary := apitypes.NewArticleSummary(articleSummary_db)

	return CreateArticleResult(articleSummary), 200, nil
}
