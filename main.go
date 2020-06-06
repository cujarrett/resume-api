package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

type Resume struct {
	Basics       Basics      `json:"basics"`
	Work         []Work      `json:"work"`
	Volunteer    []Volunteer `json:"volunteer"`
	Education    []Education `json:"education"`
	Awards       []string    `json:"awards"`
	Publications []string    `json:"publications"`
	Skills       []Skill     `json:"skills"`
	Languages    []Language  `json:"languages"`
	Interests    []Interest  `json:"interests"`
	References   []Reference `json:"references"`
}

type Basics struct {
	Name     string    `json:"name"`
	Label    string    `json:"label"`
	Picture  string    `json:"picture"`
	Email    string    `json:"email"`
	Phone    string    `json:"phone"`
	Website  string    `json:"website"`
	Summary  string    `json:"summary"`
	Location Location  `json:"location"`
	Profiles []Profile `json:"profiles"`
}

type Location struct {
	Address     string `json:"address"`
	PostalCode  string `json:"postalCode"`
	City        string `json:"city"`
	CountryCode string `json:"countryCode"`
	Region      string `json:"region"`
}

type Profile struct {
	Network  string `json:"network"`
	Username string `json:"username"`
	URL      string `json:"url"`
}

type Work struct {
	Company    string   `json:"company"`
	Position   string   `json:"position"`
	Website    string   `json:"website"`
	StartDate  string   `json:"startDate"`
	EndDate    string   `json:"endDate"`
	Summary    string   `json:"summary"`
	Highlights []string `json:"highlights"`
}

type Volunteer struct {
	Organization string   `json:"organization"`
	Position     string   `json:"position"`
	Website      string   `json:"website"`
	StartDate    string   `json:"startDate"`
	EndDate      string   `json:"endDate"`
	Summary      string   `json:"summary"`
	Highlights   []string `json:"highlights"`
}

type Education struct {
	Institution string   `json:"institution"`
	Area        string   `json:"area"`
	StudyType   string   `json:"studyType"`
	StartDate   string   `json:"startDate"`
	EndDate     string   `json:"endDate"`
	Gpa         string   `json:"gpa"`
	Courses     []string `json:"courses"`
}

type Skill struct {
	Name     string   `json:"name"`
	Level    string   `json:"level"`
	Keywords []string `json:"keywords"`
}

type Language struct {
	Language string `json:"language"`
	Fluency  string `json:"fluency"`
}

type Interest struct {
	Name     string   `json:"name"`
	Keywords []string `json:"keywords"`
}

type Reference struct {
	Name      string `json:"name"`
	Reference string `json:"reference"`
}

func getResume() *Resume {
	return &Resume{
		Basics: Basics{
			Name:    "Matt Jarrett",
			Label:   "Full Stack Developer",
			Picture: "https://mattjarrett.dev/img/header.jpg",
			Email:   "cujarrett@gmail.com",
			Phone:   "217-390-2877",
			Website: "https://mattjarrett.dev/",
			Summary: "Dad. Husband. Full Stack Developer. JavaScript Enthusiast. Aspiring Designer. Cyclist. I'm a passionate and detail oriented full stack software developer based in Dallas, Texas. Experienced leader where I enjoy design, clean code, and building positive team environments and strong developer experiences. I enjoy turning problems into solutions that make people smile. I enjoy continuous learning and exploring new technologies. When I'm not coding at home or work, I enjoy time with my family and photography.",
			Location: Location{
				Address:     "",
				PostalCode:  "",
				City:        "Dallas",
				CountryCode: "US",
				Region:      "Texas",
			},
			Profiles: []Profile{
				{
					Network:  "Linkedin",
					Username: "Matt Jarrett",
					URL:      "https://www.linkedin.com/in/matt-jarrett-303a75144/",
				}, {
					Network:  "GitHub",
					Username: "cujarrett",
					URL:      "https://github.com/cujarrett",
				}, {
					Network:  "Twitter",
					Username: "cujarrett",
					URL:      "https://twitter.com/home",
				},
			},
		},
		Work: []Work{
			{
				Company:   "State Farm",
				Position:  "Technology Engineer",
				Website:   "https://www.statefarm.com",
				StartDate: "2011-11-01",
				EndDate:   "N/A",
				Summary:   "State Farm is the largest property and casualty insurance provider in the United States. It is also the largest auto insurance provider in the United States. State Farm is ranked 36th in the 2019 Fortune 500, which lists American companies by revenue.",
				Highlights: []string{
					"Advancing Cloud Native and DevOps transformations in the telematics platform space. Working with an exciting mix of technology including public cloud, mobile, and legacy systems. Working daily on not only technology but also culture.",
					"Built new solutions using React, Node, NOSQL, REST, Kubernetes, GitLab, CI/CD, Grafana, and Prometheus for enablement of large high priority enterprise initiatives. Mentored several team members in modern technology. Designed solutions for implementation. Taught classes on modern JavaScript.",
					"Designed and implemented many successful web services allowing developers on demand access to safe fabricated data for a variety of needs and automation. It's since been used millions of times across the enterprise, enabling development teams to focus on improving customer experience.",
					"Designed and implemented an automated infrastructure solution offering a complete stand up and tear down process accomplished in minutes compared to days.",
					"Taught Java in Enterprise Java classes to spread knowledge. Active member of the college recruitment team including creation of campus events, hacks, and interviews. Mentored many college interns through a successful internships resulting in all receiving and accepting full time positions. Designed and implemented multiple solutions in Java."},
			},
		},
		Volunteer: []Volunteer{
			{
				Organization: "Tech Titans",
				Position:     "Speaker",
				Website:      "https://techtitans.org/",
				StartDate:    "2017-01-01",
				EndDate:      "N/A",
				Summary:      "Tech Titans is a forum that leverages the regional technology community to collaborate, share and inspire creative thinking that fuels tomorrow’s innovations.",
				Highlights:   []string{"I Spoke at many events over the years with audiences including teachers and students."},
			},
		},
		Education: []Education{
			{
				Institution: "Illinois State University",
				Area:        "Information Technology",
				StudyType:   "Bachelor",
				StartDate:   "2009-06-01",
				EndDate:     "2011-01-01",
				Gpa:         "3.87/4.0",
				Courses:     []string{},
			},
		},
		Awards:       []string{},
		Publications: []string{},
		Skills: []Skill{
			{
				Name:  "AWS",
				Level: "Intermediate",
				Keywords: []string{
					"Cloud", "Lambda", "S3", "CloudFront",
				},
			}, {
				Name:  "JavaScript",
				Level: "Expert",
				Keywords: []string{
					"ECMAScript", "ES6", "Node.js", "Web", "Front End",
				},
			}, {
				Name:  "React",
				Level: "Pro",
				Keywords: []string{
					"SPA", "Web", "Front End",
				},
			}, {
				Name:  "Vue",
				Level: "Intermediate",
				Keywords: []string{
					"SPA", "Web", "Front End",
				},
			}, {
				Name:  "Angular",
				Level: "Intermediate",
				Keywords: []string{
					"SPA", "Web", "Front End",
				},
			}, {
				Name:  "Go",
				Level: "Intermediate",
				Keywords: []string{
					"Golang",
				},
			}, {
				Name:  "Kubernetes",
				Level: "Pro",
				Keywords: []string{
					"K8s", "CKAD",
				},
			}, {
				Name:  "Docker",
				Level: "Intermediate",
				Keywords: []string{
					"Containers",
				},
			}, {
				Name:  "Design",
				Level: "Intermediate",
				Keywords: []string{
					"Web", "Front End",
				},
			}, {
				Name:  "Linux",
				Level: "Intermediate",
				Keywords: []string{
					"Bash", "Shell",
				},
			}, {
				Name:  "NOSQL",
				Level: "Pro",
				Keywords: []string{
					"Couch", "Dynamo",
				},
			}, {
				Name:  "SQL",
				Level: "Pro",
				Keywords: []string{
					"PostgreSQL", "Aurora",
				},
			}, {
				Name:  "GitHub",
				Level: "Expert",
				Keywords: []string{
					"SDLC", "OSS", "CI/CD",
				},
			}, {
				Name:  "GitLab",
				Level: "Expert",
				Keywords: []string{
					"SDLC", "OSS", "CI/CD",
				},
			}, {
				Name:  "Pipelines",
				Level: "Expert",
				Keywords: []string{
					"DevOps",
				},
			}, {
				Name:  "Grafana",
				Level: "Expert",
				Keywords: []string{
					"DevOps", "Observability", "Dashboards", "SRE", "SLO",
				},
			}, {
				Name:  "Prometheus",
				Level: "Expert",
				Keywords: []string{
					"DevOps", "Observability", "Dashboards", "SRE", "SLO",
				},
			}, {
				Name:     "Java",
				Level:    "Pro",
				Keywords: []string{},
			}, {
				Name:     "Photoshop",
				Level:    "Expert",
				Keywords: []string{},
			}, {
				Name:     "Python",
				Level:    "Intermediate",
				Keywords: []string{},
			},
		},
		Languages: []Language{
			{
				Language: "English",
				Fluency:  "Native speaker",
			},
		},
		Interests: []Interest{
			{
				Name: "Cycling",
				Keywords: []string{
					"MTB", "Gravel", "Road",
				},
			}, {
				Name: "Photography",
				Keywords: []string{
					"Weddings", "Landscape", "Senior",
				},
			},
		},
		References: []Reference{
			{
				Name:      "Caleb Lemoine",
				Reference: "Matt is one of the best engineers I've had the pleasure of working with and knowing personally. He would be a highly valued asset anywhere.",
			},
		},
	}
}

func handleRouting(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	if request.Method == "GET" {
		response.WriteHeader(http.StatusOK)
		responseStruct := getResume()
		bytesBuffer := new(bytes.Buffer)
		json.NewEncoder(bytesBuffer).Encode(responseStruct)

		responseBytes := bytesBuffer.Bytes()

		var prettyJSON bytes.Buffer
		error := json.Indent(&prettyJSON, responseBytes, "", "  ")
		if error != nil {
			log.Println("JSON parse error: ", error)
			return
		}
		response.Write([]byte(prettyJSON.Bytes()))
	} else {
		response.WriteHeader(http.StatusNotFound)
		response.Write([]byte(`{"message": "not found"}`))
	}
}

func main() {
	http.HandleFunc("/", handleRouting)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
