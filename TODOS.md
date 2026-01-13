# TODOs

There are plenty of things that I have done fairly haphazardly. I chose to use the built-in library for HTTP
which is good enough for a submission, but I would prefer to use Gorilla/Mux framework. 

The same goes for the GORM, I am generally not fond off the ORM frameworks, but it fits nicely into 
the "experimentation" stage. It has issues and performance is also not great when it comes down to 
using migrations. However, I could not implement my own migration library in time, so I used Gorm.

The implementation lacks tests, however not much needs to be tested and setting up a test suite for 
a database is a bit of a hassle right now. So I skipped that too. 

Tried to focus on the architecture of a service rather than focusing on the quality of the code.

Also, apologies for sloppy frontend code. Not something to be proud of but it works.