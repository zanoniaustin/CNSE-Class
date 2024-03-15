## API Design Part 2

I made two changes to the API structures to help support the use of hypermedia. The first change to the API structures is to the Vote structure. After a vote is created two links are created for that Vote. A link to the VoterId for the Voter that made the vote and this link will direct to the Voter structure so you could view information about the voter. It is similar for the other link created for the PollId used in the vote that will link to the information about the Poll voted on. I also added the same kind of link logic for the VoteValue going into the poll and then into the pollOptions to show what option was voted for.

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
             },
             {
		"href": "/poll/{PollId}/pollOptions/{VoteValue}",
		"rel": "Vote Value",
		"type" : "GET"
             }
    ]
}
```

The second change to the API structures is to the voterPoll structure. Inside the Voter's VoteHistory holds the list of voterPolls. I added in a link using the PollId to allow a link to the poll that was voted on.


```
type voterPoll struct {
	PollID    uint
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

Here was my thought process for imagining how the user goes about using these APIs. So beginning with the Votes API when viewing previous votes that have been made it is showing very simple information about the corresponding Voter, Poll, and Poll Option is associated with those votes. In order to get more information about the Voter, Poll, and Poll Option that is why I added links to take the user to the Voter, Poll, and Poll Option used in the vote so they can better see the details about that vote.

Now following the the previous line of thought if the user were to view a specific Voter's structure they would see the array of the history of Polls the Voter has voted for. It made sense to me to put in a link to those Polls to allow the user to see what Polls the Voter has voted on. I did have thoughts about adding in the VoteValue into the VoteHistory structure, but ultimately decided against it because I see the Votes API as the main storage of that sort of information, and the VoteHistory just being a sort of tracking for the Voter participating in votes.



#### Resources(Just putting down all the sites I read from as I compound all the info):
https://en.wikipedia.org/wiki/Hypermedia <br />
https://en.wikipedia.org/wiki/HATEOAS <br />
https://en.wikipedia.org/wiki/REST <br />
https://restfulapi.net/hateoas/ <br />
https://en.wikipedia.org/wiki/Richardson_Maturity_Model <br />
https://www.infoq.com/articles/webber-rest-workflow/ <br />
