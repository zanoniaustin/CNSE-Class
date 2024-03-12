
For example, the given below JSON response may be from an API like HTTP GET http://api.domain.com/management/departments/10
{
    "departmentId": 10,
    "departmentName": "Administration",
    "locationId": 1700,
    "managerId": 200,
    "links": [
        {
            "href": "10/employees",
            "rel": "employees",
            "type" : "GET"
        }
    ]
}




#### Resources(Just putting down all the sites I read from as I compound all the info):
https://en.wikipedia.org/wiki/Hypermedia <br />
https://en.wikipedia.org/wiki/HATEOAS <br />
https://en.wikipedia.org/wiki/REST <br />
https://restfulapi.net/hateoas/ <br />
https://en.wikipedia.org/wiki/Richardson_Maturity_Model <br />
https://www.infoq.com/articles/webber-rest-workflow/ <br />
