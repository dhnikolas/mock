# Self-hosted HTTP Mock Service

The self-hosted HTTP mock service typically exposes a set of endpoints that you can configure to return pre-defined responses, such as HTTP status codes, headers, and message bodies. You can use the mock service to test different scenarios, such as error responses, timeouts, or unexpected input, to ensure that your application can handle these situations correctly.
## Installation

Run with docker:

    docker run -v `pwd`:/var/mock/ -p 8111:8111 dhnikolas/mock:v2.2

## Usage

Usage with application UI:

    http://localhost:8111/ui/

Configure endpoints:

<img width="1534" alt="Снимок экрана 2023-03-30 в 02 55 50" src="https://user-images.githubusercontent.com/15860590/228693972-36d7483d-ecf5-4f21-a894-0aeb2ad5f9bf.png">

Make request and get response:

    $ curl http:/localhost:8111/api/v1/pam
    {"project_id","s98fsufhs9fusnf9si"}


Show request logs:

<img width="1688" alt="Снимок экрана 2023-03-30 в 03 01 00" src="https://user-images.githubusercontent.com/15860590/228694483-812d90a8-76e1-4a97-b1ec-3803b96107f1.png">


## Contributing

1. Fork it!
2. Create your feature branch: `git checkout -b my-new-feature`
3. Commit your changes: `git commit -am 'Add some feature'`
4. Push to the branch: `git push origin my-new-feature`
5. Submit a pull request :D

## Credits

Lead Developer - Nikolai (@dhnikolas)

## License

The MIT License (MIT)

Copyright (c) 2023 Nikolai 

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.