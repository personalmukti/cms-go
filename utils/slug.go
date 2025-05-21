package utils

import (
	"fmt"
	"regexp"
	"strings"

	"gorm.io/gorm"
)

func GenerateUniqueSlug(db *gorm.DB, title string) string {
	slug := strings.ToLower(title)
	slug = strings.ReplaceAll(slug, " ", "-")
	slug = regexp.MustCompile(`[^a-z0-9-]+`).ReplaceAllString(slug, "")

	baseSlug := slug
	counter := 1

	for {
		var count int64
		db.Table("articles").Where("slug = ?", slug).Count(&count)
		if count == 0 {
			break
		}
		slug = fmt.Sprintf("%s-%d", baseSlug, counter)
		counter++
	}

	return slug
}
