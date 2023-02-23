Feature: Behaviour of application when doing the PUT /topics/{id} endpoint, using a stripped down version of the database

    # A Background applies to all scenarios in this Feature
    Background:
        Given I have these topics:
            """
            [
                {
                    "id": "businessindustryandtrade",
                    "current": {
                        "id": "businessindustryandtrade",
                        "title": "Business, Industry and Trade",
                        "description": "Lots of information about business",
                        "state": "published",
                        "subtopics_ids": [
                            "changestobusiness",
                            "business"
                        ]
                    },
                    "next": {
                        "id": "businessindustryandtrade",
                        "title": "Business, Industry and Trade",
                        "description": "Lots of information about business",
                        "state": "published",
                        "subtopics_ids": [
                            "changestobusiness",
                            "business"
                        ]
                    }
                },
                {
                    "id": "changestobusiness",
                    "current": {
                        "id": "changestobusiness",
                        "state": "published"
                    },
                    "next": {
                        "id": "changestobusiness",
                        "state": "published"
                    }
                },
                {
                    "id": "business",
                    "current": {
                        "id": "business",
                        "state": "published"
                    },
                    "next": {
                        "id": "business",
                        "state": "published"
                    }
                }
            ]
            """
    Scenario: [Test #28] PUT /topics/businessindustryandtrade in public mode
        When I PUT "/topics/businessindustryandtrade"
            """
            {
                "title": "Business, Data and Trade",
                "description": "Lots of information about Trade",
                "release_date": "2022-10-10T08:30:00Z",
                "subtopics_ids": ["economy", "business"],
                "keywords": ["keyword1", "keyword2"],
                "state": "published"
            }
            """
        Then the HTTP status code should be "405"

    Scenario: [Test #29] Valid PUT /topics/businessindustryandtrade in private mode
        Given private endpoints are enabled
        And I am identified as "user@ons.gov.uk"
        And I am authorised

        When I PUT "/topics/businessindustryandtrade"
            """
            {
                "title": "Business, Data and Trade",
                "description": "Lots of information about Trade",
                "release_date": "2022-10-10T08:30:00Z",
                "subtopics_ids": ["changestobusiness", "economy"],
                "keywords": ["keyword1", "keyword2"],
                "state": "published"
            }
            """
        Then the HTTP status code should be "200"
        And the document in the database for id "businessindustryandtrade" should be:
            """
            {
                "id": "businessindustryandtrade",
                "title": "Business, Data and Trade",
                "description": "Lots of information about Trade",
                "keywords": ["keyword1", "keyword2"],
                "release_date": "2022-10-10T08:30:00Z",
                "state": "published",
                "subtopics_ids": [
                    "changestobusiness",
                    "economy"
                ]
            }
            """

    Scenario: [Test #30] Valid PUT /topics/businessindustryandtrade in private mode with missing fields
        Given private endpoints are enabled
        And I am identified as "user@ons.gov.uk"
        And I am authorised

        When I PUT "/topics/businessindustryandtrade"
            """
            {
                "title": "Business, Data and Trade",
                "description": "Lots of information about Trade"
            }
            """
        Then the HTTP status code should be "400"

    Scenario: [Test #31] Valid PUT /topics/businessindustryandtrade in private mode with empty strings
        Given private endpoints are enabled
        And I am identified as "user@ons.gov.uk"
        And I am authorised

        When I PUT "/topics/businessindustryandtrade"
            """
            {
                "title": "",
                "description": "",
                "release_date": "2022-10-10T08:30:00Z",
                "subtopics_ids": ["economy", "business"],
                "keywords": ["keyword1", "keyword2"],
                "state": "published"
            }
            """
        Then the HTTP status code should be "400"

    Scenario: [Test #32] Missing auth header in PUT /topics/businessindustryandtrade in private mode
        Given private endpoints are enabled
        When I PUT "/topics/businessindustryandtrade"
            """
            {
                "title": "Business, Data and Trade",
                "description": "Lots of information about Trade",
                "release_date": "2022-10-10T08:30:00Z",
                "subtopics_ids": ["economy", "business"],
                "keywords": ["keyword1", "keyword2"],
                "state": "published"
            }
            """
        Then the HTTP status code should be "401"
        

    Scenario: [Test #33] Missing request body in PUT /topics/businessindustryandtrade in private mode
        Given private endpoints are enabled
        And I am identified as "user@ons.gov.uk"
        And I am authorised
 
        When I PUT "/topics/businessindustryandtrade"
            """
            """
        Then the HTTP status code should be "400"
        And the response header "Content-Type" should be "text/plain; charset=utf-8"
        And I should receive the following response:
            """
            request body empty
            """

    Scenario: [Test #34] Invalid Topic id in PUT /topics/invalid-id in private mode
        Given private endpoints are enabled
        And I am identified as "user@ons.gov.uk"
        And I am authorised

        When I PUT "/topics/invalid-id"
            """
            {
                "title": "Business, Data and Trade",
                "description": "Lots of information about Trade",
                "release_date": "2022-10-10T08:30:00Z",
                "subtopics_ids": ["economy", "business"],
                "keywords": ["keyword1", "keyword2"],
                "state": "published"
            }
            """
        Then the HTTP status code should be "404"
        And the response header "Content-Type" should be "text/plain; charset=utf-8"
        And I should receive the following response:
            """
            topic not found
            """
