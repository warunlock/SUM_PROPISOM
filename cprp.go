package main

import (
	"fmt"
	"strconv"
)

func main() {

	fmt.Println(cprp(1234567.1, "грн.", "коп."))
	fmt.Println(cprp(1234567.1, "", ""))

}

func cprp(n float64, curr, kop string) string {

	   Nums0 := [10]string{"", "одна ", "дві ", "три ", "чотири ", "п'ять ", "шість ", "сім ", "вісім ", "дев'ять "}
	   Nums1 := [10]string{"", "один ", "два ", "три ", "чотири ", "п'ять ", "шість ", "сім ", "вісім ", "дев'ять "}
	   Nums2 := [10]string{"", "десять ", "двадцять ", "тридцять ", "сорок ", "п'ятдесят ", "шістдесят ", "сімдесят ",  "вісімдесят ", "дев'яносто "}
	   Nums3 := [10]string{"", "сто ", "двісті ", "триста ", "чотириста ", "п'ятсот ", "шістсот ", "сімсот ", "вісімсот ", "дев'ятсот "}
	   Nums4 := [10]string{"", "одна ", "дві ", "три ", "чотири ", "п'ять ", "шість ", "сім ", "вісім ", "дев'ять "}
	   Nums5 := [10]string{"десять ", "одинадцять ", "дванадцять ", "тринадцять ", "чотирнадцять ", "п'ятнадцять ", "шістнадцять ", "сімнадцять ", "вісімнадцять ", "дев'ятнадцять "}

        //розділяємо число на розряди, використовуючи допоміжну функцію Class  
	 ed     := Class(n, 1)  
	 dec    := Class(n, 2)  
	 sot    := Class(n, 3)  
	 tys    := Class(n, 4)  
	 dectys := Class(n, 5)  
	 sottys := Class(n, 6)  
	 mil    := Class(n, 7)  
	 decmil := Class(n, 8)  
	 sotmil := Class(n, 9)  
	 bil    := Class(n, 10) 
  	 bil_txt := ""
	sotmil_txt  := ""
	mil_txt :=""
	decmil_txt :=""
	sottys_txt := ""
	tys_txt := ""
	dectys_txt := ""
	ed_txt :=""
	dec_txt := ""
	sot_txt :=""


	//перевіряємо мільярди  
	switch bil  {
		case 1:
			bil_txt = Nums1[bil] + "мільярд " 
		case 2,3,4:
				
			bil_txt = Nums1[bil] + "мільяри "
		case 5,6,7,8,9:
			bil_txt = Nums1[bil] + "мільярдів "
	}

//		перевіряємо мільйони 	 
	switch sotmil  {
		case 1,2,3,4,5,6,7,8,9:
			sotmil_txt = Nums3[sotmil]
	}
  
  switch  decmil  {
   case 1:
     mil_txt = Nums5[mil] + "мільйонів "  
     goto WWW  
   case 2,3,4,5,6,7,8,9:
     decmil_txt = Nums2[decmil]
  }
   
 switch  mil { 
   case 0 : 
     if decmil > 0 { mil_txt = Nums4[mil] + "мільйонів "  }
   case 1 :
     mil_txt = Nums1[mil] + "мільйон "  
   case 2, 3, 4  :
     mil_txt = Nums1[mil] + "мільйона "  
   case 5,6,7,8,9 :
     mil_txt = Nums1[mil] + "мільйонів "  
	}
   
 if decmil == 0 && mil == 0 && sotmil != 0 { 
	sotmil_txt = sotmil_txt + "мільйонів "}

  WWW:
	sottys_txt = Nums3[sottys]

//перевіряємо тисячі  
 switch  dectys  {
   case 1:
     tys_txt = Nums5[tys] + "тисяч "  
     goto EEE  
   case 2,3,4,5,6,7,8,9:
     dectys_txt = Nums2[dectys]
}
   
 switch  tys { 
   case 0:
     if dectys > 0 {tys_txt = Nums4[tys] + "тисяч "  }
   case 1 :
     tys_txt = Nums4[tys] + "тисячa "  
   case 2, 3, 4:
     tys_txt = Nums4[tys] + "тисячі "  
   case 5,6,7,8,9 :
     tys_txt = Nums4[tys] + "тисяч "  
}
   
 if dectys == 0 && tys == 0 && sottys != 0 { sottys_txt = sottys_txt + " тисяч "  }
EEE:  
 sot_txt = Nums3[sot]


//перевіряємо десятки  
 switch dec  {
   case 1:  
     ed_txt = Nums5[ed]  
     goto RRR
   case 2,3,4,5,6,7,8,9 :
     dec_txt = Nums2[dec]
}
    
 ed_txt = Nums0[ed]  
  
RRR: 
 SUMINWORDS :=""
//формуємо підсумковий рядок  
  
 if curr == "" {
   SUMINWORDS = bil_txt + sotmil_txt + decmil_txt + mil_txt + sottys_txt + dectys_txt + tys_txt + sot_txt + dec_txt + ed_txt  
}else { 
	M_S1 := fmt.Sprintf("%.2f", n)

SUMINWORDS = bil_txt + sotmil_txt + decmil_txt + mil_txt + sottys_txt + dectys_txt + tys_txt + sot_txt + dec_txt + ed_txt + curr +" "+M_S1[len(M_S1)-2:]+" " +kop  
}   

	return SUMINWORDS
}

func Class(M float64, I int) int {
	
	
	M_S := fmt.Sprintf("%.0f", M)

	if I > len(M_S) {
		return 0
	} else 	{
		value, _ := strconv.Atoi(M_S[len(M_S)-I:len(M_S)-I+1])
		return (value)
	}


}


