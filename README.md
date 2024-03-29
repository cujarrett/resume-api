# resume-api

![logo](./media/logo.png)

<p align="center">
  <a href="https://github.com/cujarrett/resume-api/actions"><img alt="CI/CD" src="https://github.com/cujarrett/resume-api/actions/workflows/release.yml/badge.svg"></a>
</p>

## What is it?
Open source REST API for my [JSON-based standard format resume](https://jsonresume.org/). I'm using
[Go](https://golang.org/) for my app, [Terraform](https://www.terraform.io/) for my
[Infrastructure as Code (IaC)](https://en.wikipedia.org/wiki/Infrastructure_as_code) with
[Terraform Cloud](https://www.terraform.io/docs/cloud/overview.html) to provision my infrastructure,
either on demand or in response to various events. I'm deploying the app to
[AWS Lambda](https://aws.amazon.com/lambda/) with a [Circle CI](https://circleci.com/) CI/CD
pipeline.

## How can I see or use it?
You can hit the REST API via curl:
```sh
curl -L https://resume.mattjarrett.dev
```
You can view it via web browser:

https://resume.mattjarrett.dev

## Why
[To build something simple with Go to get more experience with the language](https://github.com/cujarrett/personal-goals#things-ill-do-in-2020).
To expose my resume as an API as [JSON-based standard for resumes](https://jsonresume.org/).

### Why Go?
- It's compiled. So you get small binaries.
- It's fast. (slower than c/c++ or rust) but faster than most other web programming languages.
- [Fast growing](https://octoverse.github.com/#top-languages)

## Full Write Up
I wrote a short post on how I implemented this on [Dev.to](https://dev.to/cujarrett/how-i-built-a-resume-api-w-go-terraform-and-aws-371o).
