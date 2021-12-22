# BuiilderDesignPattern

Builder Definition:
Builder is a creational design pattern, which compartmentalizes an object's initial parameters into object methods.

When to use the builder pattern?
The builder pattern is suitable when building complex objects that have similar steps but vary in specifications.



For example, let's refer to a cake factory. Imagine this cake factory...


### Subset 1
Chocolate
Vanilla
Cream Cheese
Red Velvet
Marble
Lemon
Raspberry

### Subset 2
Buttercream 
Swiss Meringue
Cream Cheese
Whipped Cream
Mascarpone
Custard
Mixed Berry

### Subset 3
None
Fondant
Message
Flowers
Design Patterns?

### Price
$X.XX


There are clearly a lot of options for each possible object variation. This could get really messy... Consider: adding more parameters, perhaps nested ones, and what if some of the cakes instances don't use those parameters?



Well the good news is the builder pattern can take care of this for us!

Builder Pattern Example:
So in the following example, there is a need for some type of ad object.

### Ad Object Attributes Needed:

Media type of the advertisement (image or video)
The advertisement's country of origin
An advertisement identification number

## Builder UML diagram
![image](https://user-images.githubusercontent.com/58267452/147111488-9a8ca68c-cb43-4e23-b5c6-510128e64cd5.png)


An advertisement identification number
