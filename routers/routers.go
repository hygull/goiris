package routers

import "github.com/kataras/iris"

import (
	"github.com/iris-contrib/middleware/basicauth"
	"goiris/conf"
	"goiris/controllers"
)

func Route() {

	/**************** Posts *******************************************/

	authentication := basicauth.Default(map[string]string{"admin": "admin@321"})

	/* To create a new post */
	iris.Post("/v1/post/new", controllers.NewPost)

	/* To display all the posts (new posts will be on top...10 posts per display) */
	iris.Post("/v1/posts/all", controllers.AllPosts)

	/* To get the posts of friends */
	iris.Post("/v1/posts/friends", controllers.FriendsPosts)

	/* To update a post */
	iris.Post("/v1/post/update", controllers.UpdateAPost)

	/* To like a post */
	iris.Post("/v1/post/like", controllers.LikeAPost)

	/* To comment a post */
	iris.Post("/v1/post/comment", controllers.CommentAPost)

	/* To share a post */
	iris.Post("/v1/post/share", controllers.ShareAPost)

	/* To delete/deactivate a post */
	iris.Post("/v1/post/deactivate", controllers.DeactivateAPost)

	/* To edit a comment */
	iris.Post("/v1/comment/edit", controllers.EditAComment)

	/* To delete/deactivate a comment */
	iris.Post("/v1/comment/deactivate", controllers.DeactivateAComment)

	/* To get all the likes related to any particular post*/
	iris.Post("/v1/post/likes", controllers.LikesOfAPost)

	/* To get all the comments related to any particular post*/
	iris.Post("/v1/post/comments", controllers.CommentsOfAPost)

	/**************** Media *******************************************/

	/* To display the list of medias eg. videos, images & texts */
	iris.Post("/v1/list/medias", controllers.MediasList)

	/* To deactivate/delete a media*/
	iris.Post("/v1/media/deactivate", controllers.DeactivateAMedia)

	/**************** User ********************************************/

	/* User login & returning JWT token */
	iris.Post("/v1/account/login", controllers.LoginIntoAccount)

	/* Updating user account */
	iris.Post("/v1/account/update", controllers.UpdateAccount)

	/* To Delete user account */
	iris.Post("/v1/account/deactivate", controllers.DeactivateAccount)

	/* To get user profile */
	iris.Post("/v1/user/profile", controllers.ProfileOfAUser)

	/* To get users medias */
	iris.Post("/v1/user/medias", controllers.MediasOfAUser)

	/* To update user address */
	iris.Post("/v1/user/address/update", controllers.UpdateAddressOfAUser)

	/* To update user experience */
	iris.Post("/v1/user/experience/update", controllers.UpdateExperienceOfAUser)

	/***************** Friend Request ********************************************/

	/* To send a friend request */
	iris.Post("/v1/request/send", controllers.SendRequest)

	/* To accept a friend request */
	iris.Post("/v1/request/accept", controllers.AcceptRequest)

	/* To reject a friend request */
	iris.Post("/v1/request/reject", controllers.RejectRequest)

	/* To unfriend a friend */
	iris.Post("/v1/unfriend", controllers.Unfriend)

	/* To get list of friend requests */
	iris.Post("/v1/requests/all", controllers.AllRequests)

	/* To search a friend */
	iris.Post("/v1/friend/search", controllers.SearchAFriend)

	/* To list suggested friends*/
	iris.Post("/v1/friends/suggested", controllers.SuggestedFriends)

	/****************** Notifications ********************************************/

	/* To get list of notifications retated to post */
	iris.Post("/v1/post/notifications", controllers.NotificationsOnAPost)

	/****************** Media upload *********************************************/

	/* To upload medias  (viedos & images) to AMAZON S3 */
	iris.Post("/v1/media/upload", controllers.UploadMedia)

	/****************** Community ************************************************/

	/* To display the list of communities */
	iris.Post("/v1/communities/list", controllers.ListCommunities)

	/****************** City offers **********************************************/

	/* To create/post/add new city offer */
	iris.Post("/v1/cityoffer/new", controllers.NewCityoffer)

	/* To update a cityoffer */
	iris.Post("/v1/cityoffer/update", controllers.UpdateCityoffer)

	/* To display cityoffers related to a particular city */
	iris.Post("/v1/cityoffers", controllers.Cityoffers)

	/* To get the list of cities */
	iris.Post("/v1/countries/list", controllers.ListCountries)

	/* To get the list of states */
	iris.Post("/v1/states/list", controllers.ListStates)

	/* To get the list of cities */
	iris.Post("/v1/cities/list", controllers.ListCities)

	/******************* Feedback ***********************************************/

	/* To send a feedback */
	iris.Post("v1/feedback/send", controllers.SendFeedback)

	/* To send a problem report */
	iris.Post("v1/problem/report", controllers.ReportAProblem)

	/* PostOffice related stuffs*/
	iris.Get("/api/v1/postoffices/", authentication, controllers.PostOffices)

	/* To display the details of only 1 postoffice */
	iris.Get("/api/v1/postoffices/:pincode", authentication, controllers.PostOffice)

	iris.Listen(":" + conf.ServerPort)
}
