# Pong online

![](https://img.shields.io/badge/made%20by-DarienMiller-blue)
![](https://img.shields.io/badge/Golang-1.20-yellow)
![](https://img.shields.io/badge/Svelte-red)
![](https://img.shields.io/badge/MongoDB-Cloud-green)
[![Netlify Status](https://api.netlify.com/api/v1/badges/fc38c13c-c8c9-4294-9701-4f9114d664fb/deploy-status)](https://app.netlify.com/sites/reliable-marshmallow-87c5cc/deploys)


## This an attempt to add online services to a web based Pong implementation.

### Built With

* [Canvas]()
* [Fiber](https://gofiber.io/)

### Requirements
* Clone the repository using `git clone https://github.com/darienmiller88/Pong-Online-Web`
* Migrate the necessary information to your local `.env` as described in the `.env_sample` file
* Run `go build` to create a root level `Messenger.exe` file, and then run `.\Messenger` to run the executable. If an executable is not needed, simply input `go run main.go` instead, or `.\fresh` to enable a server restart on change.
* `cd` into the `client` folder, and run `npm start` to the react server, which should be on port 3000

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Feel free to leave suggestions as well, I'm always looking for ways to improve!

<p align="right">(<a href="#top">back to top</a>)</p>

## License
[MIT](https://choosealicense.com/licenses/mit/)