package main

import "fmt"
import "strings"
import "strconv"



func intsliceTostring(t []int)string{


	hr:=""
	min:=""
	sec:=""
	if(t[0]<10){

		hr="0"+strconv.Itoa(t[0])
	}
	if(t[1]<10){

		min="0"+strconv.Itoa(t[1])
	}
	if(t[2]<10){

		sec="0"+strconv.Itoa(t[2])
	}

	return hr+":"+min+":"+sec
	
}

func add( t... string) string{

	curr:=[]int{00,00,00}

	for _,v:=range t{
	val:=strings.Split(v,":")

	//add the elements inside val slice
	j:=0
	for j=0;j<len(val);j++{


		res,_:=strconv.Atoi(val[j])
		curr[j]=curr[j]+res

		if curr[0]>=24{

			curr[0]=0
		}

		if curr[j]==60{

			if j>0{curr[j-1]++}

			curr[j]=00

		}
		if curr[j]>60{
			if j>0{curr[j-1]++}

			curr[j]=curr[j]-60
			
		}
		
	}
	
		
	}
	return intsliceTostring(curr)
}



func subtract(t1 string, t2 string ) string{

//subtracting t2 from t1
	T1:=strings.Split(t1,":")
	T2:=strings.Split(t2,":")

	if greater(t1,t2)==true{

		//swap the times

		for i,_:=range T1{


			T1[i],T2[i]=T2[i],T1[i]
		}
	}

	intT1:=stringsliceTOintslice(T1)
	intT2:=stringsliceTOintslice(T2)
	// fmt.Println(intT1)
	// fmt.Println(intT2)

	res:=[]int{00,00,00}

	    for i:=2;i>=0;i--{

		 if intT1[i] > intT2[i]{//take borrow from left side
			
			 fmt.Println("before",intT2)
			if intT2[i-1]==0 {
				//if 2nd place se borrow not possible
				intT2[i-2]--
				intT2[i-1]=intT2[i-1]+60  //adding into minute value
				
				
			}
			fmt.Println("after",intT2)

			intT2[i-1]--

			res[i]=intT2[i]+60-intT1[i]
			
			//fmt.Println("resultslice",res)

		 }
		
		res[i]=intT2[i]-intT1[i]
	}

	return intsliceTostring(res)

}

func stringsliceTOintslice(s[]string)[]int{

	curr:=[]int{00,00,00}
	for i,_:=range s{

		curr[i],_=strconv.Atoi(s[i])
	}

	return curr

}

func greater(t1 string,t2 string) bool{


//check whether t1 is greater than t2
	T1:=strings.Split(t1,":")
	T2:=strings.Split(t2,":")

		i:=0
		
		T1val,_:=strconv.Atoi(T1[i])
		T2val,_:=strconv.Atoi(T2[i])

		if T1val==T2val{

			
		    T1val1,_:=strconv.Atoi(T1[i+1])
			T2val2,_:=strconv.Atoi(T2[i+1])

			if T1val1==T2val2{

				
		        T1val11,_:=strconv.Atoi(T1[i+2])
				T2val22,_:=strconv.Atoi(T2[i+2])

				return T1val11>T2val22
			}
			return T1val1>T2val2

		}

		return T1val>T2val

}





func main() {

	t1 := "08:01:00"
	t2 := "07:05:05"

	//slice := add(t1, t2)
	slice:=subtract(t1,t2)

	//isgreater:=greater(t1,t2)
	

	fmt.Println(slice)

}