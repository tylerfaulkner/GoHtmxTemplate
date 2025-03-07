# Go + HTMX Template

Thank you for using this template! Th goal of this template is to simplify the setup of getting started with Go and HTMX. The primary audience for this template are developer's who want to eliminate frontend logic as much as possible.

Go was chosen as language for the server since it is fast, has a great standard library, and compiles to a ~single~ binary. HTMX is the main driver of minimzing frontend logic by extedning the HTML syntax to easily support any HTTP requests from any element and allowing the server to send HTML snippets that the HTML natively understands what to do with.

I could act all intellectual and start discussing *Hypermedia As The Engine Of Application State (HATEOAS)*, god I hope I got that right, but there are much better sources for that info.

# Getting Started

## Prerequisites

- Node [(Setup Instructions)](https://nodejs.org/en/download)
- Go [(Setup Instructions)](https://go.dev/dl/)
- Air [(Setup Instructions)](https://github.com/air-verse/air/releases/tag/v1.61.7)
- Templ CLI (`go install github.com/a-h/templ/cmd/templ@latest`)

## Running the Dev Server

1. `npm install`
2. `npm run dev`
3. Go to `http://localhost:7331`

## Creating a Production Build

1. `npm run build`