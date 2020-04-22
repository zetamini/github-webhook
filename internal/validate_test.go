package internal

import (
	"testing"
)

func TestHashing(t *testing.T) {
	// Hash
	hash := HashPayload("qweqwe", []byte(TEST_HASH_PAYLOAD))
	// Check that it matches up with what GitHub computed
	if hash != "2ab9cadffdb7fd3da51cb25666f3c15611b28151" {
		t.Errorf("Unexpected hash: %s", hash)
	}
}

// A real payload from GitHub
const TEST_HASH_PAYLOAD = `{"ref":"refs/heads/www","before":"db4a8adf35d89804875b171e796c278844e76390","after":"032e2d9fca73a73f5b2473a655c025b59c70d9d3","repository":{"id":247632243,"node_id":"MDEwOlJlcG9zaXRvcnkyNDc2MzIyNDM=","name":"barry","full_name":"yezihack/barry","private":false,"owner":{"name":"yezihack","email":"freeit@126.com","login":"yezihack","id":11530866,"node_id":"MDQ6VXNlcjExNTMwODY2","avatar_url":"https://avatars0.githubusercontent.com/u/11530866?v=4","gravatar_id":"","url":"https://api.github.com/users/yezihack","html_url":"https://github.com/yezihack","followers_url":"https://api.github.com/users/yezihack/followers","following_url":"https://api.github.com/users/yezihack/following{/other_user}","gists_url":"https://api.github.com/users/yezihack/gists{/gist_id}","starred_url":"https://api.github.com/users/yezihack/starred{/owner}{/repo}","subscriptions_url":"https://api.github.com/users/yezihack/subscriptions","organizations_url":"https://api.github.com/users/yezihack/orgs","repos_url":"https://api.github.com/users/yezihack/repos","events_url":"https://api.github.com/users/yezihack/events{/privacy}","received_events_url":"https://api.github.com/users/yezihack/received_events","type":"User","site_admin":false},"html_url":"https://github.com/yezihack/barry","description":null,"fork":false,"url":"https://github.com/yezihack/barry","forks_url":"https://api.github.com/repos/yezihack/barry/forks","keys_url":"https://api.github.com/repos/yezihack/barry/keys{/key_id}","collaborators_url":"https://api.github.com/repos/yezihack/barry/collaborators{/collaborator}","teams_url":"https://api.github.com/repos/yezihack/barry/teams","hooks_url":"https://api.github.com/repos/yezihack/barry/hooks","issue_events_url":"https://api.github.com/repos/yezihack/barry/issues/events{/number}","events_url":"https://api.github.com/repos/yezihack/barry/events","assignees_url":"https://api.github.com/repos/yezihack/barry/assignees{/user}","branches_url":"https://api.github.com/repos/yezihack/barry/branches{/branch}","tags_url":"https://api.github.com/repos/yezihack/barry/tags","blobs_url":"https://api.github.com/repos/yezihack/barry/git/blobs{/sha}","git_tags_url":"https://api.github.com/repos/yezihack/barry/git/tags{/sha}","git_refs_url":"https://api.github.com/repos/yezihack/barry/git/refs{/sha}","trees_url":"https://api.github.com/repos/yezihack/barry/git/trees{/sha}","statuses_url":"https://api.github.com/repos/yezihack/barry/statuses/{sha}","languages_url":"https://api.github.com/repos/yezihack/barry/languages","stargazers_url":"https://api.github.com/repos/yezihack/barry/stargazers","contributors_url":"https://api.github.com/repos/yezihack/barry/contributors","subscribers_url":"https://api.github.com/repos/yezihack/barry/subscribers","subscription_url":"https://api.github.com/repos/yezihack/barry/subscription","commits_url":"https://api.github.com/repos/yezihack/barry/commits{/sha}","git_commits_url":"https://api.github.com/repos/yezihack/barry/git/commits{/sha}","comments_url":"https://api.github.com/repos/yezihack/barry/comments{/number}","issue_comment_url":"https://api.github.com/repos/yezihack/barry/issues/comments{/number}","contents_url":"https://api.github.com/repos/yezihack/barry/contents/{+path}","compare_url":"https://api.github.com/repos/yezihack/barry/compare/{base}...{head}","merges_url":"https://api.github.com/repos/yezihack/barry/merges","archive_url":"https://api.github.com/repos/yezihack/barry/{archive_format}{/ref}","downloads_url":"https://api.github.com/repos/yezihack/barry/downloads","issues_url":"https://api.github.com/repos/yezihack/barry/issues{/number}","pulls_url":"https://api.github.com/repos/yezihack/barry/pulls{/number}","milestones_url":"https://api.github.com/repos/yezihack/barry/milestones{/number}","notifications_url":"https://api.github.com/repos/yezihack/barry/notifications{?since,all,participating}","labels_url":"https://api.github.com/repos/yezihack/barry/labels{/name}","releases_url":"https://api.github.com/repos/yezihack/barry/releases{/id}","deployments_url":"https://api.github.com/repos/yezihack/barry/deployments","created_at":1584341864,"updated_at":"2020-03-16T07:41:35Z","pushed_at":1587543535,"git_url":"git://github.com/yezihack/barry.git","ssh_url":"git@github.com:yezihack/barry.git","clone_url":"https://github.com/yezihack/barry.git","svn_url":"https://github.com/yezihack/barry","homepage":null,"size":2547,"stargazers_count":0,"watchers_count":0,"language":null,"has_issues":true,"has_projects":true,"has_downloads":true,"has_wiki":true,"has_pages":false,"forks_count":0,"mirror_url":null,"archived":false,"disabled":false,"open_issues_count":0,"license":null,"forks":0,"open_issues":0,"watchers":0,"default_branch":"master","stargazers":0,"master_branch":"master"},"pusher":{"name":"yezihack","email":"freeit@126.com"},"sender":{"login":"yezihack","id":11530866,"node_id":"MDQ6VXNlcjExNTMwODY2","avatar_url":"https://avatars0.githubusercontent.com/u/11530866?v=4","gravatar_id":"","url":"https://api.github.com/users/yezihack","html_url":"https://github.com/yezihack","followers_url":"https://api.github.com/users/yezihack/followers","following_url":"https://api.github.com/users/yezihack/following{/other_user}","gists_url":"https://api.github.com/users/yezihack/gists{/gist_id}","starred_url":"https://api.github.com/users/yezihack/starred{/owner}{/repo}","subscriptions_url":"https://api.github.com/users/yezihack/subscriptions","organizations_url":"https://api.github.com/users/yezihack/orgs","repos_url":"https://api.github.com/users/yezihack/repos","events_url":"https://api.github.com/users/yezihack/events{/privacy}","received_events_url":"https://api.github.com/users/yezihack/received_events","type":"User","site_admin":false},"created":false,"deleted":false,"forced":false,"base_ref":null,"compare":"https://github.com/yezihack/barry/compare/db4a8adf35d8...032e2d9fca73","commits":[{"id":"032e2d9fca73a73f5b2473a655c025b59c70d9d3","tree_id":"af8a5182c175db2b7b3214ab531840e10ec4b630","distinct":true,"message":"readme","timestamp":"2020-04-22T16:18:46+08:00","url":"https://github.com/yezihack/barry/commit/032e2d9fca73a73f5b2473a655c025b59c70d9d3","author":{"name":"barry","email":"freeit@126.com","username":"yezihack"},"committer":{"name":"barry","email":"freeit@126.com","username":"yezihack"},"added":[],"removed":[],"modified":["README.md"]}],"head_commit":{"id":"032e2d9fca73a73f5b2473a655c025b59c70d9d3","tree_id":"af8a5182c175db2b7b3214ab531840e10ec4b630","distinct":true,"message":"readme","timestamp":"2020-04-22T16:18:46+08:00","url":"https://github.com/yezihack/barry/commit/032e2d9fca73a73f5b2473a655c025b59c70d9d3","author":{"name":"barry","email":"freeit@126.com","username":"yezihack"},"committer":{"name":"barry","email":"freeit@126.com","username":"yezihack"},"added":[],"removed":[],"modified":["README.md"]}}`