package xlab
package main

imoprt {
	"context"
	"fmt"
	"log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/bson"
}
var db *mongo.Database
var todolist *mongo.Collection
type todo struct{
	id int
	name string
}
const url="mongodb://user:pass@sample.host:27017/?maxPoolSize=20&w=majority"
func initMongdb(){
	option:=options.Client().ApplyURI(url)
	client,err:=mongo,Connect(context.TODO(),option)
	if err!=nil{
		log.Fatal(err)
	}
	err=client.Ping(context.TODO,nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Print("connect mongo success")
	db = client.Database("test")
	stuColl = db.Collection("ToDo")
}
func Insert(thing todo){//增
	one,err:=todolist.Insert(context.TODO,thing)
}
func Delete(filter interface{}){//删
	one,err:=todolist.Delete(context.TODO(),filter)
}
func Query(filter interface{})todo {//查
	var result todo
	err:=todolist.FindOne(context.TODO(),filter).Decode(&res)
	return result
}
func Update(filter interface{},condition interface{}){//改
	one,err:=todolist.UpdateOne(context.TODO(),filter,condition)
}
func main()  {

}