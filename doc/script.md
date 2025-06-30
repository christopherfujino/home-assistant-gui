# Planttadore

<link rel="stylesheet" href="https://www.christopherfujino.com/global-style.min.css">

Started with this post: https://system76.com/weekend-project/house-plant-monitoring/

Tech stack:

- Raspberry Pi 4
- 800 x 480 LCD screen
- [Home Assistant](https://www.home-assistant.io/)
- Bluetooth Plant sensor
- Flower Care mobile app
- Raylib game library
- raylib-go bindings for the Go programming language


## Script

- Initial idea
    - Combine
        - Lexie's interest in plants
        - Lexie's skills in designing new products
        - My interest in writing new kinds of software
    - Create a device that can monitor a plant
- Philosophy
    - Should not feel like a "computer", but rather a device with a single purpose
    - I love technology, but not hyperconnectivity
        - Should be self-contained, no networking, no mobile app
        - To find out how the plant is doing, you still need to walk over to it
    - Should be open source
- Technical problems to solve:
    - Sensor
        - Obtaining (easy)
        - Integrating it into the system (hard?)
    - Custom software
        - Integrating with sensor data (hard?)
        - Presenting a user interface (medium easy)
    - Display
        - Obtaining (easy)
        - Integrating it (easy)
- Home Assistant OS is headless
- Tech stack slide deck
- Screen was upside down
- Latency between sensor and device
- Heat
    - Throttling FPS
- What you'll need:
    - Raspberry Pi 4
    - 800 x 480 LCD screen
    - [Home Assistant](https://www.home-assistant.io/)
    - Bluetooth Plant sensor
    - Flower Care mobile app
    - Raylib game library
    - raylib-go bindings for the Go programming language
