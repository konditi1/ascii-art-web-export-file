# ASCII Art Generator

Welcome to the ASCII Art Generator project! This web application allows users to input text and instantly convert it into ASCII art.

## Features

- **Text Input:** Enter any text you want to convert into ASCII art.
- **ASCII Art Generation:** Converts the entered text into ASCII art using selected banner styles.
- **Responsive Design:** The application works seamlessly on desktops, tablets, and mobile devices.

## Technologies Used

- HTML5
- CSS3 (with responsive design principles)
- Go (for handling HTTP requests, responses, and ASCII art generation)

## Installation

Clone the repository:

```sh
git clone https://github.com/jesee-kuya/ascii-art-web
cd ascii-art-web
```

## Usage

1. **Run the Program:**
   ```sh
   go run .
   ```
   This command starts the ASCII art generator server.

2. **Open the Application:**
   Navigate to `localhost:7050/` in your web browser.

3. **Input Text:**
   Enter the desired text in the provided input field.

4. **Select a Banner:**
   Choose a banner file to style your ASCII art.

5. **Generate ASCII Art:**
   Click on the "Generate" button to create ASCII art based on your input and banner selection.

6. **Copy or Download:**
   Use the "Copy" button to copy the generated ASCII art to your clipboard, or download it as needed.

## Algorithm

The application utilizes HTML templates for the user interface, which communicates with a Go-based server for processing. Here's a simplified breakdown of the process:

- **User Interaction:** HTML form collects user input (text and banner choice).
- **Server-side Processing:** Go server receives the input, retrieves the selected banner file, and generates ASCII art based on the provided text.
- **Response:** ASCII art is returned as a string and displayed on the user interface.

This workflow ensures efficient and seamless conversion of text to ASCII art, tailored to user preferences.

## Authors

- [Jesee Kuya](https://github.com/jesee-kuya)
- [James Muchiri](https://github.com/j1mmy7z7)
- [Fena Onditi](https://github.com/konditi1)

---