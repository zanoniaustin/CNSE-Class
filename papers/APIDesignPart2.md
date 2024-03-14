## API Design Part 2

The first change to the API structures is to the Vote structure. After a vote is created two links are created for that Vote. A link to the VoterId for the Voter that "made" the voteand this link will direct to the Voter structure so you could view information about the voter. It is similar for the other link created for the PollId used in the vote that will link to the information about the Poll voted on.

```
type Vote struct {
	VoteID    uint
	VoterID   uint
	PollID    uint
	VoteValue uint
	"links": [
             {
		"href": "/voter/{VoterId}",
		"rel": "Voter",
		"type" : "GET"
             },
             {
		"href": "/poll/{PollId}",
		"rel": "Poll",
		"type" : "GET"
             }
    ]
}
```

The second change to the API structures is to the voterPoll structure. Inside the Voter's VoteHistory holds the list of voterPolls. Now I don't know if it was just missed 


```
type voterPoll struct {
	PollID    uint
	VoteValue uint
	VoteDate  time.Time
	"links": [
	     {
		"href": "/poll/{PollId}",
		"rel": "Poll",
		"type" : "GET"
	     }
	]
}

type Voter struct {
	VoterID     uint
	FirstName   string
	LastName    string
	VoteHistory []voterPoll
}
```




#### Resources(Just putting down all the sites I read from as I compound all the info):
https://en.wikipedia.org/wiki/Hypermedia <br />
https://en.wikipedia.org/wiki/HATEOAS <br />
https://en.wikipedia.org/wiki/REST <br />
https://restfulapi.net/hateoas/ <br />
https://en.wikipedia.org/wiki/Richardson_Maturity_Model <br />
https://www.infoq.com/articles/webber-rest-workflow/ <br />
