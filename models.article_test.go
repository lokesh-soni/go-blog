// models.article_test.go

package main

import "testing"

// Test the function that fetches all articles
func TestGetAllArticles(t *testing.T) {
  alist := getAllArticles()

  // Check that the length of the list of articles returned is the
  // same as the length of the global variable holding the list
  if len(alist) != len(articleList) {
    t.Fail()
  }

  // Check that each member is identical
  for i, v := range alist {
    if v.Content != articleList[i].Content ||
      v.ID != articleList[i].ID ||
      v.Title != articleList[i].Title {

      t.Fail()
      break
    }
  }
}

// Test the function that fetche an Article by its ID
func TestGetArticleByID(t *testing.T) {
  a, err := getArticleByID(1)

  if err != nil || a.ID != 1 || a.Title != "Article 1" || a.Content != "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum" {
    t.Fail()
  }
}