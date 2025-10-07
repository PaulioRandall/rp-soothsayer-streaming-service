# Rapid Prototype: Soothsayer

I'm trying to create a tool called _Soothsayer_ for supporting thematic-based analysis of video content. Because the video content is nearly always sensitive data, I aim to create a desktop based application where everything is local to the machine. The early focus of the app will be usability. I want to make it functional first to avoid designing something that cannot be implemented. 

I plan to use a web-based frontend if possible as this would give me access to all the frontend libraries and knowledge for creating a user-centered application. I will only explore other solutions, e.g. Java JWT or C++ desktop apps, if it's not feasible.

However, I'd rather have the backend in a nice fast language like [Go](https://go.dev/), hence I'm strongly looking at [Wails](https://wails.io/) as the framework. However, only a fool would marry themselves to a solution. I will look at alternatives once I've exhausted preferred possibilites.

## Aims

*Aim:* To make a decision about how I'm going to proceed with the video streaming aspect of the project.

*Objective 1:* To explore technical barriers and solutions to a desktop application where video content is streamed from backend to a web-based frontend.

*Objective 2:* To understand how much effort would be needed to develop such a solution.

## Results

Explored several approaches.

### Event-based streaming with Wails

I quickly found [jaesung9507/playgo](https://github.com/jaesung9507/playgo). The repo is an example of how to stream content with [Wails](https://wails.io/) from a [Go](https://go.dev/) backend to a web frontend using events.

It works but video quality is dynamic and poor on first load for a local file. However, it indicates it's possible to use a none server-based approach! I.e. stream files to a web frontend within the boundaries of the app.

The main problem is that I only understand half of what is going on in the repo. I would need to spend time reading and playing with video encoding, transmission, and chunking to apply the solution in the context of my problem.

### Loading video files

I tried to create a simple solution where the video is loaded by the frontend, entirely dispensing with the backend for streaming content. While this works for [Firefox](https://www.firefox.com/en-GB/?redirect_source=mozilla-org) or [Chromium](https://www.chromium.org/getting-involved/download-chromium/), it doesn't seem to work for HTML views such as webpage viewers that [Wails](https://wails.io/) and other frameworks actually use. This is annoying but understandable.

I didn't look into:
- Configuring the webview tools to allow local file loading. However, the solution needs to work for Linux, Windows, and Mac.

### Locally hosted server

Finally, I created a locally hosted server that serves files from a test folder. Works well and quite simple. It's also very fast with a 320MB video loading instantly. The solution only took me a few hours, including testing and experimentation.

However, I'm not sure how secure locally hosted stuff is. I'm pretty sure one would need to expose ports and stuff for the service to be accessible externally. However, any locally running application would be able to scan sockets and probably find the hosted service. Then again, if you're running locally you've probably already got access to the files.

I did a very simple static file serve because that was easiest. But I think the frontend loads the whole video into memory before playing. While this is great for playabck and seeking performance, it may be an issue for low spec machines. I expect most users would can deal with 2-3 hour video footage providing they have a 8GB of ram (most interviews being analysed are less than an hour so usually less than 400MB). 

To add a layer of security:
- The application could generate a token when loaded and passed to the frontend. Each backend request will reject requests without the token.

To deal with potential performance issues:
- The backend could implement adaptive streaming to avoid high memory usage, e.g. [HLS](https://en.wikipedia.org/wiki/HTTP_Live_Streaming). However, this would require converting video files for delivery and learning how to implement a HLS.

## Conclusions

I'm going to go with the locally hosted server solution for now. It's very simple, quick, and low effort. I can always add CSRF style tokens to add a layer of security. I can also implement [HLS](https://en.wikipedia.org/wiki/HTTP_Live_Streaming) end points later if memory is an issue.

I can revist the video streaming approach when I've got a fully functional MVP. What's most important is the UI. It needs to be optimised for quick and low effort qualitative data coding, unlike [Nvivo](https://lumivero.com/products/nvivo/). The video streaming can evolve later.
