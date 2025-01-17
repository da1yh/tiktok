package test

import (
	"net/http"
	"testing"
)

func TestFeed(t *testing.T) {
	e := newExpect(t)
	feedResp := e.GET("/douyin/feed/").Expect().Status(http.StatusOK).JSON().Object()
	feedResp.Value("status_code").Number().IsEqual(0)
	feedResp.Value("video_list").Array().Length().Gt(0)
	for _, element := range feedResp.Value("video_list").Array().Iter() {
		video := element.Object()
		video.ContainsKey("id")
		video.ContainsKey("author")
		video.Value("play_url").String().NotEmpty()
		video.Value("cover_url").String().NotEmpty()
	}
}

func TestUserAction(t *testing.T) {
	e := newExpect(t)
	userName, userPassword := "s1mple", "7355608"
	registerResp := e.POST("/douyin/user/register/").WithQuery("username", userName).WithQuery("password", userPassword).
		WithFormField("username", userName).WithFormField("password", userPassword).
		Expect().Status(http.StatusOK).JSON().Object()
	registerResp.Value("status_code").Number().IsEqual(0)
	registerResp.Value("user_id").Number().Gt(0)
	registerResp.Value("token").String().Length().Gt(0)

	loginResp := e.POST("/douyin/user/login/").WithQuery("username", userName).WithQuery("password", userPassword).
		WithFormField("username", userName).WithFormField("password", userPassword).
		Expect().Status(http.StatusOK).JSON().Object()
	loginResp.Value("status_code").Number().IsEqual(0)
	loginResp.Value("user_id").Number().Gt(0)
	loginResp.Value("token").String().Length().Gt(0)

	userId, token := int64(loginResp.Value("user_id").Number().Raw()), loginResp.Value("token").String().Raw()
	userResp := e.GET("/douyin/user/").WithQuery("user_id", userId).WithQuery("token", token).
		Expect().Status(http.StatusOK).JSON().Object()

	userResp.Value("status_code").Number().IsEqual(0)
	userInfo := userResp.Value("user").Object()
	userInfo.NotEmpty()
	userInfo.Value("id").Number().Gt(0)
	userInfo.Value("name").String().Length().Gt(0)
}

func TestPublishAction(t *testing.T) {
	e := newExpect(t)
	userId, token := getTestUser("Niko", e)
	publishResp := e.POST("/douyin/publish/action/").
		WithMultipart().WithFormField("token", token).WithFormField("title", "bear").
		WithFile("data", "../public/bear.mp4").
		Expect().Status(http.StatusOK).JSON().Object()
	publishResp.Value("status_code").Number().IsEqual(0)

	//for testing publish_list, check the length of video list is greater than 0,
	// for every video in video list, there exist necessary keys and url is not empty

	publishListResp := e.GET("/douyin/publish/list").
		WithQuery("user_id", userId).WithQuery("token", token).
		Expect().Status(http.StatusOK).JSON().Object()
	publishListResp.Value("status_code").Number().IsEqual(0)
	publishListResp.Value("video_list").Array().Length().Gt(0)
	for _, elememt := range publishListResp.Value("video_list").Array().Iter() {
		video := elememt.Object()
		video.ContainsKey("id")
		video.ContainsKey("author")
		video.Value("play_url").String().NotEmpty()
		video.Value("cover_url").String().NotEmpty()
	}

}
