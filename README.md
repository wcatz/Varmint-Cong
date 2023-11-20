![Bill Murray](https://raw.githubusercontent.com/wcatz/Varmint-Cong/master/static/media/bill-gopher.webp)
# # Varmint-Cong
Go front end development with Gorilla, Htmx, tailwinds and Air.

## How to run

[Download Tailwind CSS CLI](https://tailwindcss.com/blog/standalone-cli)

Add it to $PATH

sudo mv tailwindscss /usr/local/bin/

install [Air](https://github.com/cosmtrek/air)

go install github.com/cosmtrek/air@latest

Clone this repo and cd into it

Run the code with air and watch and build css with tailwinds cli

air && tailwindcss -m -i static/css/main.css -o static/css/output.min.css --watch

server is on localhost:8080

refresh page to see changes.