# In-Memory CRUD Project

## Overview

This project is a simple CRUD (Create, Read, Update, Delete) application that **does not use a database**.  
Instead, all data is stored and managed **in the RAM (memory)** during the runtime of the application.

## Features

- Create new records
- Read and view existing records
- Update existing records
- Delete records
- Data is kept in memory (RAM) and will be reset every time the server/application restarts

## How It Works

- When you run the application, it initializes an empty list/array in memory.
- All create operations add new items into this list.
- Read operations retrieve data from the list.
- Update operations modify the existing items in the list.
- Delete operations remove items from the list.
- Since there is no database, **all data will be lost when the application stops**.

## Quick Start

```bash
git clone https://github.com/your-username/in-memory-crud-project.git
cd in-memory-crud-project
go mod tidy
go run main.go
```

You can then access the application at `http://localhost:your-port`.

## Example Post Request

{
  "Title" : "Your Title",
  "Body" : "Your Body"
}
