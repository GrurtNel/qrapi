package post

import (
	"gopkg.in/mgo.v2/bson"
)

type PostCategoryJoiner struct {
	Post         `bson:",inline"`
	CategoryName string `bson:"category_name" json:"category_name"`
	AuthorName   string `bson:"author_name" json:"author_name"`
	Thumb        string `json:"thumb"`
}

func GetPosts(userID string) ([]*PostCategoryJoiner, error) {
	var match = bson.M{
		"dtime": bson.M{
			"$ne": 0,
		},
	}
	if userID != "" {
		match["author"] = userID
	}
	var pipe = queryPostCategory()
	pipe = append([]bson.M{
		bson.M{
			"$match": match,
		},
	}, pipe...)
	pipe = append(pipe, bson.M{"$sort": bson.M{"ctime": -1}})
	var posts []*PostCategoryJoiner
	err := postTable.Pipe(pipe).All(&posts)
	return posts, err
}

func GetPost(postID string) (*PostCategoryJoiner, error) {
	var post *PostCategoryJoiner
	var match = bson.M{
		"_id": postID,
	}
	var pipe = queryPostCategory()
	pipe = append([]bson.M{
		bson.M{
			"$match": match,
		},
	}, pipe...)
	err := postTable.Pipe(pipe).One(&post)
	return post, err
}

const LIMIT = 10

func GetAllPosts(catID string, page int64) ([]*PostCategoryJoiner, error) {
	var pipe = queryPostCategory()
	if catID != "" {
		var match = bson.M{
			"$match": bson.M{
				"category": catID,
				"dtime": bson.M{
					"$ne": 0,
				},
			},
		}
		pipe = append([]bson.M{match}, pipe...)
	}
	pipe = append(pipe, bson.M{"$skip": (page - 1) * LIMIT})
	pipe = append(pipe, bson.M{"$limit": LIMIT})
	pipe = append(pipe, bson.M{"$sort": bson.M{"ctime": -1}})
	var posts []*PostCategoryJoiner
	err := postTable.Pipe(pipe).All(&posts)
	return posts, err
}
func queryPostCategory() []bson.M {
	var joinCategory = bson.M{
		"$lookup": bson.M{
			"from":         "category",
			"localField":   "category",
			"foreignField": "_id",
			"as":           "category_join",
		},
	}

	var joinAuthor = bson.M{
		"$lookup": bson.M{
			"from":         "user",
			"localField":   "author",
			"foreignField": "_id",
			"as":           "author_join",
		},
	}

	var unwindCategory = bson.M{"$unwind": "$category_join"}
	var unwindAuthor = bson.M{"$unwind": "$author_join"}
	var replaceRoot = bson.M{
		"$replaceRoot": bson.M{
			"newRoot": bson.M{
				"$mergeObjects": []interface{}{
					bson.M{
						"category_name": "$category_join.name",
						"author_name":   "$author_join.name",
						"thumb": bson.M{
							"$concat": []string{staticPath, "", "$_id"},
						},
					},
					"$$ROOT",
				},
			},
		},
	}
	return []bson.M{joinCategory, joinAuthor, unwindCategory, unwindAuthor, replaceRoot}
}

const staticPath = "http://163.44.206.108:3006/static/post/"

func (p *PostCategoryJoiner) TransformThumb(staticPath string) {
	p.Thumb = staticPath + p.ID
}
