## API Design Part 2
```
type voterPoll struct {
	PollID    uint
	VoteValue uint
	VoteDate  time.Time
	"links": [
		{
			"href": "10/employees",
			"rel": "employees",
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

```
type Vote struct {
	VoteID    uint
	VoterID   uint
	PollID    uint
	VoteValue uint
	"links": [
        {
            "href": "10/employees",
            "rel": "employees",
            "type" : "GET"
        },
        {
            "href": "10/employees",
            "rel": "employees",
            "type" : "GET"
        }
    ]
}
```




#### Resources(Just putting down all the sites I read from as I compound all the info):
https://en.wikipedia.org/wiki/Hypermedia <br />
https://en.wikipedia.org/wiki/HATEOAS <br />
https://en.wikipedia.org/wiki/REST <br />
https://restfulapi.net/hateoas/ <br />
https://en.wikipedia.org/wiki/Richardson_Maturity_Model <br />
https://www.infoq.com/articles/webber-rest-workflow/ <br />
