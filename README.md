
# Undisclosed Prior Golang Interview Take Home

I took this test for an undisclosed company, but I feel the questions and assignment were interesting enough that I want it in my github as an example of my though process. Thanks!

## q&a

The following sections contain questions that you should answer inline in this README.md file. Below each question, insert your response using tasteful Markdown choices.

### tooling

Provide your perference for each of the categories below.

- operating system?

Good question. I guess I prefer MacOS because that's what everyone uses, but I feel more at home in a flavor of the linux environment - Arch is supposedly for the cool kids these days, but Ubuntu brings me back to college. Not a fan of walled gardens. I like a lot of things that Microsoft has come out with - but not Windows.

- text editor?

VSCode! Speaking of things I like about Microsoft, VSCode is definitely a big one. I think I just like how modifiable it is - although I have to say it can have a large memory footprint given that it is written in Electron (code_helper can eat up system resources over time - maybe it has a memory leak? idk). Other editors that I like are other "micro" like editors like Atom or Sublime. I find things like Eclipse and Netbeans, .NET to a lesser extent, to be poorly written in my opinion - they are large and complicated but I don't see how this does much to increase productivity. Some editors for mobile, like Swift and Android Studio (Android, again, to a lesser extent), rely on a GUI for screen creation and also allow it to be done programmatically - this I think is dangerous (2 sources of truth), so that squigs me out a little. As far as EMACS/VIM/NANO - I like nano for editor editing, but I haven't really found the use case of hotkeying several commands and working entirely in the terminal. Again YMMV.

- terminal type?

I usually just use vanilla terminal in MacOS or Linux. I don't really use terminals from third parties, although I am using a program called Hyper which provides nice colors for fonts and such. Ideally I think the terminal should be as close to a "source of truth" you can make so you know at least that whatever program you run through terminal isn't crashing because terminal is crashing. Powershell / Windows terminal is, to put it bluntly, not good. 

- scripting languauge?

Depends on the use case, but I really like both Python and Javascript. Javascript because it allows manipulating the dom and both Javascript and Python because they are so ubiquitous that someone's probably solved your problem before if it's not too unique. I think they get a "bad rap" because they are both simple and powerful which has led to an explosion of code (some of which is a higher quality than others). But I think these are both positives - they are productive, popular langauges that have large communities.

### pros & cons

Provide some salient pros and cons for each of the scenarios described below.

- windows vs. linux

By and large these days the GUI and user experience is largely the same - they both use windows to hold data, have start bar navigation (for the most part on linux - I'm using Ubuntu as "linux"), they both have file explorers etc. Where they differ is that Windows is proprietary and linux is open source (setting aside Redhat which is an enterprise distro), Windows bills itself as easier to use and relies less on the command line and more on GUI commands. Linux, being open source, relies on a community to provide patches, bug releases, hardware driver support, etc. Hardware support can sometimes be especially lacking - if you have a unique wifi or sound card, for example - as you may have to find a patched driver to support it. Windows on the other hand costs money and they put advertisements in the start bar, and you can't look at the source code so there are probably other nefarious things in there - but! - it's easier to use and works out of the box. Except for powershell (mostly kidding). Also if you want to work with .NET stack (although Core just came out for Mac, it's not yet on Linux and I don't think Core necessarily has everything yet) or some sort of Windows GUI software, or anything else Windows native, you need to use Windows. Overall, I would prefer linux.

- mono-repository vs. multi-repository

Depends on the size and complexity really. If you have an app with many small tightly coupled sub-parts and the app itself is small you should probably stick with a mono-repository. Overall mono-repositories reduce complexity so no engineer is going to get confused in what code goes in what repo, that one repo must be pulled before another or the build queue doesn't work, etc. On the other hand if you have a large application with a few loosely coupled sub-parts (maybe even purely functional sub-parts!) it may make sense to put these parts in another repository. Sometimes the distinction between parts can be "this is our frontend code, this is our backend code". Then again, Google just uses a mono-repo for almost ALL their code (<a href='http://delivery.acm.org/10.1145/2860000/2854146/p78-potvin.pdf?ip=67.198.78.26&id=2854146&acc=OA&key=4D4702B0C3E38B35%2E4D4702B0C3E38B35%2E4D4702B0C3E38B35%2E5945DC2EABF3343C&__acm__=1546530284_67093cadb5f87ccdcaac91a7ef38a4ee'>BORG</a>). (Interesting tidbit on drawbacks: Drawbacks include having to create and scale tools for development and execution and maintain code health, as well as potential for codebase complexity (such as unnecessary dependencies)). Overall it depends on a judgement call on which method to take and by how much.

- strongly typed vs. untyped languages

Strongly typed languages force variables to take type declarations and untyped languages interpret variable types. I've read that most scientific studies of programmer productivity have shown that the number of errors per line of code appears more or less constant among languages - so adding types to a language probably increases verbosity and therefore errors. Probably the developer experience to coding typed languages can feel "slower" to get to compile. On the other hand...probably most developers prefer more strongly typed languages. While type coercion is sometimes not possible, and you often have to do frustrating things like declare the types of variables that functions take, those frustrating things can mean the difference between compile time errors (annoying but predictable) and runtime errors (often "bayesian"). I don't know how I feel honestly - I'm pretty agnostic on types.

- monolith vs. microservice architecture

A monolith means the code base is one big piece with no independently working sub-pieces, whereas a microservice is many small code bases that talk to each other by passing messages (this explanation is at a 30,000 foot level). The worst case scenario of a monolith is you have a code base that is large and complex so any one developer may have difficulty understanding it, it is typically written in an "enterprise language" like Java or .NET for stability reasons or to talk to legacy systems, and if it breaks *where* it is breaking (or is slow) can be hard to determine. On the other hand, microservices can become rats nests of spaghetti code where state is not well held and race conditions are common or hard to debug. The happy balance is to probably have a CRUD backend, some sane way to manage state, and micro-services where they are needed but not "just because". Again, this is a matter of some discrimination.


### git

Answer the following questions regarding git.

- Who created git?

Linus Torvalds (https://www.google.com/search?client=firefox-b-1-ab&q=Who+created+git%3F)

- Describe your preferred git branching model

![courtesy of `https://medium.com/@grazibonizi/the-best-branching-model-to-work-with-git-4008a8098e6a`](https://cdn-images-1.medium.com/max/1455/1*2YagIpX6LuauC3ASpwHekg.png)

In the above I take master to be "production" ready code to be shipped, and integration to be the main development branch. Features can be assigned as tickets or for larger code bases those features could be broken out among developers. A much deeper tree structure than maybe that one more layer is probably not advisable, and you should switch to multiple repos or **really** think about a highly scalable model.

- What is a git remote?

From the perspective of a local machine a `remote branchName` implies that `branchName` is a branch on a remote server or github. 

- What is git LFS?

**Git Large File Storage (LFS) replaces large files such as audio samples, videos, datasets, and graphics with text pointers inside Git, while storing the file contents on a remote server like GitHub.com or GitHub Enterprise.** You can read more here <a href="https://git-lfs.github.com/">https://git-lfs.github.com/</a>.

- Provide the command you would use to stage, commit, and push local changes

I typically do the following, but it depends on the workflow. 

  0) *(OPTIONAL)* `git pull origin devbranch` (pull `devbranch` if you are making a feature off of it)
  1) `git add .` (adds current directory to staging area ready for commit)
  2) `git commit -m "commit message goes here"` (commit your staging area to be pushed)
  3) *(OPTIONAL)* `git merge devbranch` (merge in `devbranch` if there is possibility remote branch has change since pulling, then make sure to squash git conflicts)
  4) `git push origin mybranch` or *OPTIONAL* `git push origin devbranch`  (push to your own branch or the devbranch)

- What does `git config` do?

It's a way of configuring the git program (changing font colors and user email addresses among other things). You can read more here <a href="https://www.atlassian.com/git/tutorials/setting-up-a-repository/git-config"></a>.

- What does `git stash` do?

**Use git stash when you want to record the current state of the working directory and the index, but want to go back to a clean working directory. The command saves your local modifications away and reverts the working directory to match the HEAD commit.** For more information see <a href="https://git-scm.com/docs/git-stash">here</a>

- Is the `master` branch of a repository special?

**The “master” branch in Git is not a special branch. It is exactly like any other branch. The only reason nearly every repository has one is that the git init command creates it by default and most people don’t bother to change it.** For more information see <a href="https://git-scm.com/book/en/v2/Git-Branching-Branches-in-a-Nutshell">here</a>.

### docker

Answer the following questions regarding Docker.

- What is Docker?

Docker is a "lightweight" VM that uses a configuration file in your code to wrap your codebase in an (usually linux based) operating system of your choice. This sandboxes your code so that your code can run and ship on the VM OS with its own dependencies and information can't get in or out unless explicitly allowed (write once, run everywhere). Its main appeal is its incredibly small CPU overhead to run and ease of developer use.

- How would you compare Docker versus a virtual machine?

This is a pretty good post <a href="https://stackoverflow.com/questions/16047306/how-is-docker-different-from-a-virtual-machine">here</a>. This is a seemingly simple question that can get incredibly complicated quickly if you need to dive under the hood of how Docker works, but the short answer is 1) it's lighter weight 2) requires less system resources because they are shared. This may mean that sometimes there are things you can't do in Docker because you need a full VM. For most things that run on a linux distro I haven't run into this problem, but I wouldn't try to run iOS or Windows as a Docker container but you could as a VM. I would say that if you want to do something esoteric in a sandboxed environment and you're finding that Docker is having issues looking at whether a VM is necessary is probably a good step.

- What is a Docker image?

I can't really improve on this explanation <a href="https://stackoverflow.com/questions/23735149/what-is-the-difference-between-a-docker-image-and-a-container">here</a>. The answer with pictures is probably clearest.

- What is a Docker container?

I can't really improve on this explanation <a href="https://stackoverflow.com/questions/23735149/what-is-the-difference-between-a-docker-image-and-a-container">here</a>. The answer with pictures is probably clearest.

- What is Docker Hub?

Literally like github but for Docker images. You can make an account and everything. It's pretty neat!

- How do you create a Docker container?

First you use `docker build...` to take your Dockerfile and build an image. Then you use `docker run...` to take that image and build a container from it which will immediately spin up and be accessible.

- Do I lose my data when the Docker container exits?

Maybe? If you've persisted your data using Volumes then no, or if you've written to an outside file, the internet, or your own database then also no. Otherwise probably yes.

- What, in your opinion, is the most exciting potential use for Docker?

So, I subscribe to the "smallest pipe theory". Imagine you have a series of connected pipes, with their widths varying in size. Then, given some constant water pressure where water flows from one end to another, only by increasing the size of the smallest pipe will the flow rate increase. So, for a long time shipping code easily was a smallest pipe, and now it's not. We now have a glut of computing power out there that is potentially unused and we could containerize. Silly example - but we could have household appliances that could take new OS/code by wifi updates - and cheaply! But if I were to say what I would really want to see it would be Docker containers self perpetuating and creating more of themselves like a giant <a href="http://pmav.eu/stuff/javascript-game-of-life-v3.1.1/">conway's game of life</a>. We could even imagine antiviral software that moved through a company's systems and had a reproduction rate tied to how many viruses they found, as a far future speculation. 

- What is `docker-compose`?

A way to run multiple docker containers simultaneously and get them to talk to each other. A happy medium between a simple dockerfile, but before you have a real need for kubernetes.

### javascript

Answer the following questions regarding JavaScript.

- What are your thoughts on JavaScript?

It's VHS to **insert your pet language**'s Betamax. It was never to designed to become as popular as it has, but probably because of this it has an incredible simplicity of use. Untyped, interpreted it doesn't require much more to get started than a javascript file and a server (maybe some html). For all the flack it gets I like it - I'm a big fan of having things be only as complicated as they need to be and for what Javascript's use case is (managing the DOM, some light data manipulation and networking) I don't want to have to manage garbage collection or types etc. So I like it I guess.

- What's the last cool thing you implemented in JavaScript?

I was making a blog at zennify.me for a while. A recentish project I'm working is a sort of collage art project writtin in next.js (server side react essentially). It's difficult to describe, but if you've ever seen the MTV show Daria, it's a bit how I think `Sick Sad World` would look as a website. It uses liberal use of `https://deepdreamgenerator.com/` which is a cool website that allows you to take the style of one picture and use it to paint the content of another using Google's Deep Dream AI program they came out with a couple years ago. Coupling that with GIMP (a photo editing tool) allows me to make some pretty cool static hologram effects. This is definitely an art project and done only for my own happiness and is therefore a little "fashion forward". Currently in development, here is a screenshot: 

![Psych](readmeIMG/jsProj.png)

- Why do you think node.js is so popular?

Couple of reasons. It's written in javascript, and since you basically have to know how to use that for the frontend it means frontend devs can start building backend servers out of the box. It's async by default which is super nice. It uses the NPM registry which I've personally heard is the reason for a few shops in town to switch from .NET which has a dying eco-system. It does all this while being "fast enough" for networking and web dev, so it hits the sweet spot of ease of use and efficient engineering.

- What's the difference between `npm` and `npx`?

NPX does a <a href="https://blog.npmjs.org/post/162869356040/introducing-npx-an-npm-package-runner">lot of things</a>. Basically it's a utility that's build on top of npm that allows execution of packages without global installation. I think it was released shortly after a scare where malicious packages were added to the npm registry (say misspelled package names by on letter containing a virus etc), and the fix was rolled into a larger utility. Using npx is generally preferable to using npm where applicable.

- Should you commit `node_modules` to git?

No those should be git ignored and then make sure in the install directions to include an `npm install` command.

### go

Answer the following questions regarding Go.

- What are your thoughts on Go?

It's pretty amazing. I wrote a blog post about it which can be found on zennify.me, but probably the best summation is this 20 minute youtube <a href='https://www.youtube.com/watch?time_continue=4&v=rFejpH_tAHM'>video</a> given by one of the creators of the language Rob Pike. The main takeaway is that Go is designed to have few ways of doing as many things as possible so there are less orthogonal implementations leading to code complexity. One of the big ways you can see this in the language is that it uses structs but almost no other data types. It's a bare bones but highly performant language that is designed primarily for systems networking. I don't think it will ever be as ubiquitious as say Python, because it's made to be simple and simply isn't as extensible, but it does have the backing of Google so it'll be around for a while at least.

- What's the last cool thing you implemented in Go?

So I was working on a fun thing and I wanted authenticated log in. So first I got username and password stored and retrieved in bcrypt, but then I realized I needed a way to log them in once of course. So what I ended up doing was giving a person a JWT when they signed in and a cookie that was valid for some amount of time. Then if they came back and still had the cookie I auto logged them in. I also think I had it so that if they were active in the last few minutes their cookie would roll over (I think I may have just assigned them a new valid one). It ended up being a lot of work, but I might go down a different session strategy if I were to do that again.

- Say you build a Go app on your mac. Can you copy that app to a linux box and run it?

Not if it's compiled, but you can <a href="https://www.yellowduck.be/posts/cross-compile/">cross compile it</a>. If you want to run it without compilation you would have to install Go and change your paths to coincide whats in your bash file. Or you can use docker like I've done in `go/src/example1`.

- What's the difference between `go build` and `go install`?

**Both go install and go build will compile the package in the current directory when ran without additional arguments. If the package is package main, go build will place the resulting executable in the current directory. go install will put the executable in $GOPATH/bin (using the first element of $GOPATH, if you have more than one). If the package is not package main, go install and go build will compile the package, displaying any errors encountered** You can read more <a href='https://pocketgophers.com/go-install-vs-go-build/'>here</a>.

- Explain `GOPATH`

`GOPATH` is where your bash_profile tells golang to look for dependencies. All of your projects are intended to exist under this `GOPATH` in the `src` folder. There is also a `GOROOT` which are system level dependencies. The whole thing has been a bit of a headache as it's made git commits harder and generally I think they are moving away from `GOPATH` (see below).

- How do you deal with Go dependencies?
 
Honestly, I use a Dockerfile. It solves the dependency issue, but I also don't have to worry about GOPATH or any of that either, which is pretty great. I have an example of how I use Docker to do this in `go/example1`, but I think it's a common hack. As of the latest version of Go, 1.11, <a href="https://ukiahsmith.com/blog/a-gentle-introduction-to-golang-modules/">modules</a> look like the `package.json` equivalent that will handle dependencies. It's pretty new though so I haven't had a chance to play with it.

- What's special about `internal` packages, if anything?

Internal packages are a way of scoping packages so they can only be accessed within a defined tree structure and not outside of it. This means you can have a set of packages that can communicate with each other, but can't be imported by packages that are not within the tree. You can learn more <a href="https://golang.org/doc/go1.4#internalpackages">here</a>. 

- Explain Go's Java Virtual Machine equivalent?

Technically it doesn't have one. See <a href="https://golang.org/doc/faq#runtime">here</a>.

- What's a closure?

A closure is a fancy way of saying that we return the value of a variable from a function so we can escape normal cascade hoisting rules. A simple example is <a href='https://tour.golang.org/moretypes/25'> here </a>. In this case we are "closing" the variable `sum` so that it's accessible in the for loop *even though* sum is declared in a function that is **not** a parent or grandparent of main. You can find another good example and a longer explanation <a href='https://stackoverflow.com/questions/36636/what-is-a-closure'>here</a>. I don't use them often as they seem overly complicated for what they do, but it's a good pattern to maintain variables and only allow certain functions write to them. 

- What's a goroutine? How is it different than a standard OS thread?

Goroutines can be thought of "lightweight" threads - meaning you can treat them as threads when you think about concurrency (barring very heavy usage). Go multiplexes (fancy word for load balances more or less) goroutines over a few threads and opens more when threads start becoming blocked. You can read more in this good <a href="https://golangbot.com/goroutines/">article</a>.

- What's a channel?

A channel is a way for goroutines to communicate with each other even though they are in separate "threads" and therefore can finish their operations asynchronously.

- When would you prefer to use traditional mutexes?

Basically if you have a bunch of goroutines all manipulating the same variable then you need to make sure that they have exclusive access to the variable and are not overwritten by other goroutines. So a mutex allows a goroutine to lock access to the state while it's manipulating it, and unlock access once done for the next goroutine. Here is a very good <a href="https://gobyexample.com/mutexes">example</a>.

- What's the difference between arrays and slices?

An array is a fixed size data type and a slice is a mutable size reference to an array. Arrays are annoying to pass to functions as you have to pass a pointer, so unless specifically looking to use an array you usually use a slice - it's the most like a vector one is used to using in any other standard data manipulation language. The internals can get <a href="https://appliedgo.net/slices/">a little hairy</a>, but as a rule of thumb it's best to be aware that there are interesting differences if errors start appearing near some slices or arrays.

- Explain the purpose and use case of `context`

```type Context interface {
    Deadline() (deadline time.Time, ok bool)
    Done() <-chan struct{}
    Err() error
    Value(key interface{}) interface{}
```

👆 that's the entire context interface. It's used basically for closing goroutines or defering behavior of goroutines by some interval. It's like a concurrent `setTimeout` or `setInterval` in Javascript. I think I've used them once or twice, but it's not a tool I reach for often.

- Why use `gofmt`?

`gofmt` standardizes your code so it's easier for other developers to read. Moving spaces around and such.

- What's an interface? Should you pass it by value?

An interface is a type that you give to an object; this type is special - it has a list of methods with the values those methods take in and the values those methods return. Those methods must be defined somewhere. But this now allows you to call those functions on the value we gave the interface type. It's a little hard to put into words but here is a good <a href="https://gobyexample.com/interfaces">example</a>. As for passing by value, I'm not sure I understand this question - I don't think interfaces have a value in the traditional sense, but I wouldn't pass it by reference, I would just pass the interface. 

- gRPC vs. REST? Talk about pros and cons for each.

So REST is older, more stable, and a lot of systems rely on it. But it relies on HTTP1.1 which is a single threaded non-socket channel with a high system overhead. gRPC is the new shiny and it works with HTTP2 which supports much higher throughput with a smaller overhead, multiplexing, data streaming all sorts of stuff. It also tries to fix the REST problems of systems becoming spaghetti code over time by enforcing a pattern (configuration over convention) - personally this is billed as a plus, but patterns always work until they don't and then you're stuck. Probably this is the tool to reach for if you really need HTTP2 and are willing to eat the technical learning. 

- When does `row.Close()` get executed in the following snippet:

```go
func() {
  for {
    row, err := db.Query("SELECT ...")
    if err != nil {
      panic(err)
    }
    defer row.Close()
  }
}
```

The cursor (which is what we call the memory the pointer that is searching the db) needs to be freed up but only when it either returns from its search successfully or throws an error. So that's when `defer row.Close()` happens.

### kafka

- What is Kafka?

Kafka is an event streaming queue designed to take in events from multiple different sources (often microservices) and stream them back. It has a lot of functionality out of the box, and is sort of like a massive pub-sub system on steriods.

- How is Kafka different than RabbitMQ?

There are a lot of differences. But I would say that the main one is that Kafka is better at managing a timeline of state, which RabbitMQ lacks, while RabbitMQ is a bit simpler at small scale. You can read a quick summary <a href="https://stackoverflow.com/questions/42151544/is-there-any-reason-to-use-rabbitmq-over-kafka">here</a>.

- What's the difference between exactly-once and at-least-once semantics?

In the first case a message is delivered once and it requires that state be handled on the other hand to make sure that it is not delivered again, while in the second all the guarantee that is required by the delivering server is a message from the client sayin that at least one message was delivered.

- What's the difference between a topic and partition?

In metaphorical terms, a topic is a "vector" of messages that can be read from and written to by clients with access. A partition is used in the same sense as a hard drive partition is used. A topic can be split into multiple partitions and spread across a kafka cluster, so that, of course, no partition is greater in size than the server that hosts it and all the partitions of the topic are accessible via the cluster. The kafka cluster uses its magic under the hood to route client requests to the correct partition when they request a topic. It's a distribued system.


## implementations

The following sections asks you to implement mini-projects. You should add your project files directly to this repository.

### docker

Create a directory called `docker`. Now implement the following within this directory:

1. Create a `Dockerfile` so that an image can be built that runs `htop`. Your image should work against the following docker commands:

    `$ docker run -it --rm <YOUR_IMAGE>` 

    Should run `htop`.

    `$ docker run -it --rm <YOUR_IMAGE> --version`

    Should echo the version of `htop` installed.

Answer can be found in `docker/runDocker.sh`.
  
2. Create a `docker-compose.yml` so that you can bring up two services at once that run `htop`. In other words, running the following should create two `htop` instances:

    `$ cd ./docker && docker-compose up`

Answer can be found in `docker/runDockerCompose.sh`.

### go

Create a directory called `go` and assume this will be your `GOPATH`. Now implement the following within this directory:

Create an http server that implements the following API specifications. Note, you do not need to implement this API with a persistance layer; data can simply live in memory.

Assume this service will exist as a microservice. Make sure your code is testable.

Below I talked a bit about looking up the Swagger API and finding the answer and my thoughts on that. I wanted to also share how I would approach the problem. The code exists in the `/go/src/example1` folder and all you have to do is `run.sh`. 

The entry point of the application is in `main.go`, which routes incoming requests to `/requests`. From there, if need be, `/controller` is called for data manipulation, otherwiser `/responses` is called directly. Global state is managed in `/state` and is called by the other folders as needed to pull or push state. 

I've shown in `/requests` how you can break packages up, so there are two files there as an example, and I included a simple unit test in the `/controller` but have not tested everything. (For funsies the unit test failing breaks the Dockerfile from compiling which is good practice).

---

I looked up the code for this and found the answer. I learned a bit about Swagger and Open API which was fun - I will say that the code generation using <a href="https://github.com/go-swagger/go-swagger">go-swagger</a> was **much** different and more verbose than what this example is and it doesn't appear that swagger natively works with golang yet (it claims to on its website, but its command line tool spits out errors if you try). If you would like to see the route I was going down with goswagger check it out <a href="https://goswagger.io/tutorial/custom-server.html">here</a> although it definitely seemed broken - I was using the below provided swagger doc to generate the code, but again it was very verbose. I'm curious to know if this code was code generated or not - swagger/openAPI certainly an interesting technology, although I would worry that overly relying upon code generation will have the same issues that lisp does (write once, read never). Looking into the repos for open api (<a href="https://github.com/go-openapi/runtime">ie example</a>) leave me slightly worried. I've opened a documentation issue, so maybe someone will point me in the right direction if there's docs elsewhere.

Here are a couple alternatives to swagger that I think are pretty interesting. Personally I've heard the performance of buffalo is not *the best* but there appears to be much better documentation and I like it a bit more than swagger I think: <a href='https://gobuffalo.io/en'>linkylink</a>. Also, <a href='http://www.gorillatoolkit.org/'>gorilla!</a>

To run the code go into the `GOPATH` directory and `go get` and then `go run main.go` in `src`. Make doubly sure that you use `/api` appended to your path to get the correct api returns (ie `http://127.0.0.1:8344/api/pets`).

```swagger
swagger: "2.0"
info:
  version: 1.0.0
  title: Swagger Petstore
  license:
    name: MIT
host: petstore.swagger.io
basePath: /v1
schemes:
  - http
consumes:
  - application/json
produces:
  - application/json
paths:
  /pets:
    get:
      summary: List all pets
      operationId: listPets
      tags:
        - pets
      parameters:
        - name: limit
          in: query
          description: How many items to return at one time (max 100)
          required: false
          type: integer
          format: int32
      responses:
        "200":
          description: A paged array of pets
          headers:
            x-next:
              type: string
              description: A link to the next page of responses
          schema:
            $ref: '#/definitions/Pets'
        default:
          description: unexpected error
          schema:
            $ref: '#/definitions/Error'
    post:
      summary: Create a pet
      operationId: createPets
      tags:
        - pets
      responses:
        "201":
          description: Null response
        default:
          description: unexpected error
          schema:
            $ref: '#/definitions/Error'
  /pets/{petId}:
    get:
      summary: Info for a specific pet
      operationId: showPetById
      tags:
        - pets
      parameters:
        - name: petId
          in: path
          required: true
          description: The id of the pet to retrieve
          type: string
      responses:
        "200":
          description: Expected response to a valid request
          schema:
            $ref: '#/definitions/Pets'
        default:
          description: unexpected error
          schema:
            $ref: '#/definitions/Error'
definitions:
  Pet:
    required:
      - id
      - name
    properties:
      id:
        type: integer
        format: int64
      name:
        type: string
      tag:
        type: string
  Pets:
    type: array
    items:
      $ref: '#/definitions/Pet'
  Error:
    required:
      - code
      - message
    properties:
      code:
        type: integer
        format: int32
      message:
        type: string
```






