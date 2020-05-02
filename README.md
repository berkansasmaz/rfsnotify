# R-FSNotify
<p>

  <img alt="build" src="https://github.com/berkansasmaz/rfsnotify/workflows/Go/badge.svg" /> 
 <img alt="Version" src="https://img.shields.io/badge/version-0.6.5-blue.svg?cacheSeconds=2592000" />
   <a href="https://github.com/berkansasmaz/Ketum/blob/master/LICENSE" target="_blank">
    <img alt="License: GPL-3.0" src="https://img.shields.io/badge/License-GPL3.0-yellow.svg" />
		<a href="https://twitter.com/berkansasmazz" target="_blank">
    <img alt="Twitter: berkansasmazz" src="https://img.shields.io/twitter/follow/berkansasmazz.svg?style=social" />
  </a>
  </a>
</p>

> This project is still under development and its API is not stable. Please use it with caution or wait until this message is removed and the first version is released.

A recursive file watcher package based on `github.com/fsnotify/fsnotify`.

Unfortunately, `fsnotify` does not have recursive watching capability, and you need to write your way to find all the sub-folders and files underneath and add them by using its `Add()` method.

R-FsNotify is the solution for that problem. It automatically watches all of your files under a directory.

### Warning
It is important to realize that this package is only for recursive file watcher. You can disable the recursive nature of this package. Hence, you can use github.com/fsnotify/fsnotify if you need a non-recursive watcher instead. 

Since this package is still under development, the API surface may change as the new requirements come up. Therefore, until the first release is fully published, use this library with caution. <b>[Feel free to contact me!](mailto:berkansasmazz@gmail.com?subject=Ketum&body=Hi,)</b>

## Unit Tests
This project is covered by various unit tests in the `rfsnotify_test.go` file. My intention is to keep it <strong> 100% covered </strong>. In case you would like to create PR for this project, please make sure that your code does not reduce the test coverage score. However, I am aware of the fact that not everything can be unit-testable, but it is still a crucial practice to keep unit tests in mind while contributing to this project.
  
  <img alt="unit-tests" src="https://user-images.githubusercontent.com/31216880/76708386-6a0eac80-6707-11ea-978e-c6ab625df22f.png" /> 
  
## 🤝 Contributing

Contributions, issues and feature requests are welcome!<br/>Feel free to check [issues page](https://github.com/berkansasmaz/rfsnotify/issues). You can also take a look at the [contributing guide](https://dev.to/janessatran/a-beginner-s-guide-to-contributing-to-open-source-4fen).

## Show your support

Give a ⭐️ if this project helped you!

## 📝 License

Copyright © 2020 [Berkan](https://github.com/berkansasmaz).<br />
This project is [GPL3.O](https://github.com/berkansasmaz/rfsnotify/blob/master/LICENSE) licensed.
