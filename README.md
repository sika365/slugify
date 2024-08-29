# slugifyr

A Go package to convert strings into slug format, supporting UTF-8, Persian, and English characters.

## Features

- Converts any input string to a URL-friendly slug.
- Supports UTF-8 characters, including Persian, English, and other languages.
- Removes special characters and replaces spaces with hyphens (`-`).
- Handles mixed-language inputs correctly.

## Installation

To install the package, run:

```bash
go get github.com/sika365/slugify
```

## Usage

Here is how to use the slugifyr package:

```go
package main

import (
	"fmt"
	"github.com/sika365/slugify"
)

func main() {
	slug, err := slugifyr.Slugify("Hello سلام")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(slug) // Output: hello-سلام
}

```

## Testing

To run tests, use:

```bash
go test
```

## License

This project is licensed under the MIT License.

# اسلاگفایر (slugifyr)

یک بسته‌ی Go برای تبدیل رشته‌ها به قالب اسلاگ (slug)، با پشتیبانی از کاراکترهای UTF-8، فارسی و انگلیسی.

## ویژگی‌ها

- تبدیل هر رشته ورودی به یک اسلاگ مناسب برای URL.
- پشتیبانی از کاراکترهای UTF-8، شامل فارسی، انگلیسی و دیگر زبان‌ها.
- حذف کاراکترهای ویژه و جایگزینی فاصله‌ها با خط تیره (-).
- مدیریت صحیح ورودی‌های ترکیبی از زبان‌های مختلف.

## نصب

برای نصب بسته، این دستور را اجرا کنید:

```
go get github.com/sika365/slugify
```

## استفاده

برای استفاده از بسته slugifyr، به این شکل عمل کنید:

```go
package main

import (
	"fmt"
	"github.com/sika365/slugify"
)

func main() {
	slug, err := slugifyr.Slugify("Hello سلام")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(slug) // خروجی: hello-سلام
}
```

## تست

برای اجرای تست‌ها از این دستور استفاده کنید:

```bash
go test
```

## مجوز

### Explanation of the Content

- **Introduction (English and Persian):** Briefly describes what the package does.
- **Features (English and Persian):** Lists the main features.
- **Installation (English and Persian):** Explains how to install the package.
- **Usage (English and Persian):** Provides a sample code snippet to demonstrate usage.
- **Testing (English and Persian):** Shows how to run the tests.
- **License (English and Persian):** Mentions the type of license for the package.

Make sure to replace `fibonachyy` with your actual GitHub username or repository name.
