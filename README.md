## learn design pattern Iterator pattern
```
concept :iterator is to traverse a container 
and access every container's element
```
- build data structure
> `usertype/usertype.go`
- use data structure to build a container
#### interface
- `collection`(collection/collection.go) > aggregate data 
```
+ userGroup []usertype.UserType
- createIterator() iterator
```

- iterator
```
+ Index int
+ userGroup []usertype.UserType
- HasNext() bool
- GetNext() *uesrtype.UserType
```

#### concrete 

- userCollection
```
+ users []usertype.User
- createIterator() iterator
> return &userIterator.UserIterator{ Users: usercollection.Users}
```

- userIterator  
```
+index int, User usertype.User
HasNext() bool : check if this browse is terminated
GetNext() *usertype.User : if HasNext === true,
move to next `node(user)` and index +1
else terminate (return nil)
```
