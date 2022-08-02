## Exercise

When I was a grad student at Purdue University, we used to use several different software packages for example for VLSI design or C++ programming running on Sun/HP workstations. 
And they were licensed for the use of limited number of concurrent users. When the concurrent number of users reached the limit for a software package then new users couldn't run it at all and were notified that all of the licenses have all been taken up and retry later. 

Now here is the **problem**: 

-How can you implement such a license mechanism so that no change to either in software package or client machines would be necessary. 
Design and implement your solution using proxy design pattern. 

-Implement your own error mechanism to represent the case where no license is available.