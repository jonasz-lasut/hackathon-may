# github.com/jonasz-lasut/hackathon-may

Welcome to May hackathon generated docs

## Routes

<details>
<summary>`/article`</summary>

- [RequestID]()
- [Logger]()
- [Recoverer]()
- [URLFormat]()
- [SetContentType.func1]()
- **/article**
	- **/**
		- _GET_
			- [onasz-lasut/hackathon-may/server.DatabaseHandler.ArticleListGetter-fm]()
		- _PUT_
			- [onasz-lasut/hackathon-may/server.DatabaseHandler.ArticleCreator-fm]()

</details>
<details>
<summary>`/article/remove/{articleID}`</summary>

- [RequestID]()
- [Logger]()
- [Recoverer]()
- [URLFormat]()
- [SetContentType.func1]()
- **/article/remove/{articleID}**
	- [adminOnly]()
	- **/**
		- _DELETE_
			- [onasz-lasut/hackathon-may/server.DatabaseHandler.ArticleDeleter-fm]()

</details>
<details>
<summary>`/article/{articleID}`</summary>

- [RequestID]()
- [Logger]()
- [Recoverer]()
- [URLFormat]()
- [SetContentType.func1]()
- **/article/{articleID}**
	- **/**
		- _GET_
			- [onasz-lasut/hackathon-may/server.DatabaseHandler.ArticleGetter-fm]()
		- _POST_
			- [onasz-lasut/hackathon-may/server.DatabaseHandler.ArticleUpdater-fm]()

</details>
<details>
<summary>`/healthz`</summary>

- [RequestID]()
- [Logger]()
- [Recoverer]()
- [URLFormat]()
- [SetContentType.func1]()
- **/healthz**
	- _GET_
		- [onasz-lasut/hackathon-may/server.DatabaseHandler.HealthcheckHandler-fm]()

</details>

Total # of routes: 4
