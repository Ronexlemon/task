package config

import (
	"context"
	"managementapi/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoColletions struct {
	mongoColletion *mongo.Collection
}

func (m *MongoColletions) CreateTask(t *model.Task) (interface{}, error) {
	result, err := m.mongoColletion.InsertOne(context.Background(), t)

	if err != nil {
		return nil, err
	}
	return result, nil

}

func (m *MongoColletions) AllTask()(interface{},error){
	result, err := m.mongoColletion.Find(context.Background(),bson.D{})
	if err != nil {
		return nil, err
		}
		var tasks []model.Task
		if err := result.All(context.Background(), &tasks); err != nil {
			return nil, err
			}
		return tasks, nil
}


func (m *MongoColletions) DeleteTask(task_id string)(int64, error){
	result,err := m.mongoColletion.DeleteOne(context.Background(),bson.D{{Key: "task_id", Value: task_id}})
	if err != nil {
		return 0, err
		}
		return result.DeletedCount, nil
}

func (m *MongoColletions) UpdateTaskToOngoing(task_id string,status *model.Task)(int64,error){
	result,err := m.mongoColletion.UpdateOne(context.Background(),bson.D{{Key: "task_id",Value: task_id}},bson.D{{Key: "$set",Value: status.Status.Ongoing}})
	if err != nil {
		return 0, err
		}
		return result.ModifiedCount, nil
}

