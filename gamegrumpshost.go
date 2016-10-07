package main

import (
	"fmt"
	"net/http"
)

var html string = `<!DOCTYPE html>
<html><head><meta http-equiv="Content-Type" content="text/html; charset=windows-1252"><style>
body {
background-color: #000;
}
div.youtube {
top: 50%;
left: 50%;
}
</style>

<title>Game Groomples</title>
</head>

<body onresize="itResizedOhShitKneeGrow()">


<center><div id="youtubeContainer"></div></center>

<script>
var urls = ["vSzz4c_E0Bw","qNqfYtd3HTg","8fRi3Wwo3A8","hPAADCAOKpc","DO0FjaR2Z-8","BQI2vK31N6Q","gKYLKUoVohI","FQfHeBva7hA","Ir1YLfXqsKI","AiTKbO1h7Qw","75d9_WftjK0","Fjrm2gLuHmo","U39QXgCRoa8","cjWSxJLmu_s","rqIZR6gfT_I","W1ElAY59DLU","-pqY7b-mNE0","6oVIJq_Hm5M","VuUJItINOSk","Dd6fVoq1-7E","Hjp9YGId1MU","yQ2VDxUNpGQ","RXGu7Sm-Qag","mvUKXTp3nWY","SY481N-K79w","37IYpCjyOkU","PaU8ResFbMw","ONSKAsxjf7o","qrs7JUFXTvk","mWiYoR0pX7k","snuYz1DzCQA","N9zSpaHz-3E","O-GMdgI3vhA","P9vUIH5Rv8M","yQSkj7WDJ7U","X5KveAh6dQw","ljt0nIt4Ywk","iw69uRquih8","bpuulWuxK5c","wntOFj7Sj1U","f6V05lHuFJw","aIOsNP1Q_rA","h0TlGLe1LGA","Lpvd7AenCKc","2gkTPrhDwaU","9l1QBLRwJ3s","uXW11B4dEqU","48yW_7ro6iA","k6l__bsdsBY","CiHJrzjs0eU","rxbm052_TXE","7XJ4VBn4Uww","OMYJa808i_Q","qmnyp7TLs-Y","JhvRgimkINM","E1KCuPpCb_g","PevJ-ID3NMc","9w_sFSA6oVQ","XgVsXI1riNc","uQ2qHN69msI","ThB5SLndtYI","JgwxuusiH2k","4Xc3HxLlkqE","I8uEAvdVbnM","EYk93hW_Wyc","Wz7DN6Q9WsI","9Us3KTlEBJE","X5iZUqPUsPU","_l5RxRdlzrQ","ottzIFrdYjE","KR5DQL6f_Cc","mXB9eZoXYZI","Vmq5G1c1RvA","xR4H9r61PsA","2uIn8cfpECA","HQ01O4lxbIc","Ibhtaz3H1Xs","1oZn37qrU6Y","BFQ8ktrqKQU","XgloH5S7xuk","u3P0q6x9zgk","8avBDI3b98I","Ok1XvdU61rM","5SKCCi-stiE","IJtpTrMT5j0","P_6RRkv5QPQ","GGjWtdbLHBo","sHKTwwkYol8","Wyzcv6j60K0","bksAf_2rAi0","kbuXozAsyTc","BFvcGC1_cOo","7SnLRPI7sWY","BsqDOYOIHeE","5Y8Ck28mgWI","dcALdk9ZdjI","4J09Hnenjjc","x--cCnmCitE","MQJRtq-a2O0","M4Qq1uswxQY","cY_P0g1wO2k","KrReyA4Zsvs","pDVqDGOpDo4","WoJP-sszVJY","YyOdPwA1H54","sqaSlVMhoe0","DRMsyCpUdkk","9pvBS_YV9Us","mkBaz2IMiMU","VJJ8knazdXY","uVUmZ9VB2II","lYvLZT6BLR4","Wa_WTPTJ7cQ","W2GyW74eBp4","ig3InKjjKok","rdwE46fB3Uk","l5eztItHUxs","CORtX1JJ0Pc","2i7xK2M2i8A","e_gPZGiV_m8","uiofcRg6XEo","cUx4_xFglgQ","I25evqGoF3o","CESsQ_M5xJI","jswzoBYq7EI","sD_m_OVKHKY","LTkRTERhL-c","5YOZuBzY0sE","eFM_3zTVo_E","zPPR8xu90sc","A8RscVmsdk8","2liTFdfBJu0","lCgWuN3Kesk","RNKNBEP0L3A","74yK4-_3K0E","E8Egrmm0GeE","NTqUvmG6Bdg","8U9x8AGIR5A","Da7z-TjJmOs","4wSB3AsSDyA","cH7x-DwaVu8","do-kG-UpEsY","eD842Vl6a5o","XNgW2h5yQG4","GJdYC0X7o08","hdP_BImG4XY","DI48SIr899Y","yKr56PO4GSM","mvG39NmgVPk","r1PBtHe4unE","8WuuAVAskvo","9FgEF8T8vYc","8TsTdF7vkhk","WXt3oFfbgCw","NZYkpM3-mqw","IYl8otozuh0","omv53CWafdM","OY5vWfB8_1M","4MIh40E8Trk","qXj7E6dFPwo","ZuK7l7cepA8","FZcE_4KuR_I","JyWfTcfZPkU","_4NhM5gahds","giA1VLk-O_A","ViBzl1eQT0A","1lWsIhsfrH8","mShHHdMTI6k","X19FZ4HYXZE","JlISiyMJyyc","zfTnx7iVhVo","faH6HPEiCaw","_N5P79-_9oc","XiIYl1oG4rQ","uBaj9qfEnaI","bdybqXHn2qQ","1UWOTX9uPCw","Omkw_GiNvMY","94-WH4a_WWA","hb0eYtx8iyg","dJoDPAaM1KM","pU-2Cb439rE","naO4xBIbzeg","z0zMvcv3S48","Vc7m3G3XEKU","xQMyc95jjNg","WaDw9TslRxc","pPIrvAX1WXY","gJuxS0qEw90","pjhdxX80orI","alckNXBWPno","4RS9gSkJC6Q","APnmMb-WRzA","i-GBPRThak8","gkOmQ_jkvB8","CVF-OfN7Kzs","Q6UFdwjrsss","sixukZDKiXk","B50VuLQyR20","YpPD82rrzos","nsr43CLD3yc","YDGrmEMpWTo","WViZ8_UP6-U","zE5zNI2_d2s","2S5n6PooJck","D4VNvihRRWM","aRuIE6Rplbs","Hy3EzaZZz-U","m0w7ERsmZGo","gNUD3b3pPLA","Myi9dPEV2Xs","YEaE4VcB-7k","kxF9B2V_uW8","8uJTv3HDOsw","j_rImC4opME","lK62_5_qspg","5FC48f-x7OE","RzL8diuv84M","xyKHx-uast0","1qpBlOCgWQk","7kRYzOOCHvY","Z3zEVPttE8w","ms-9bmKmivE","t7AtQHXCW5s","l1kL3jSu6mY","uVTaF4s9xfI","p60J1s0tDZk","ohw_5GHvDYw","Aiv85IScHqs","zct-Ufi5Lsg","AdeLpndCux8","7Gs482ay94A","QDTiPAx8Egw","oRNIAD1UtUk","t1AuPrdDYzQ","owlbbOm1E4A","s7GUiVnxa6s","QWv5WWEw82k","DTNc2g4yo18","BWqLdKl_NBk","CokFZgPEMoU","pBpcW6-YwMY","WuGsRmEEx6s","odIIO-GslDE","-2phajBXE8g","e52HcmapB4M","5EtbvH80C0A","7DstPlYYk-M","Tl_yUW1FnkY","SSt442tv4hM","gyj_uuyQXHU","nEthu8FaiX8","iGQJmZBENcA","b3ZZ6nnTtWw","ifTY8MF8EeE","JyK6tb5HbTs","k-tVN3TXaSk","yvPU99Yr4oI","OzE0LWd5SmM","NwIickc_a6Q","Su34O7WBt2U","knbTwjorw6E","bE9ubKn-jV0","SHyO-cR9ugY","EqwPzoKAC7s","hwAjdTXPqfE","uVUL6I8D2eU","soRnAvIUrw8","HJifVqZAfZQ","MFIyTF8G0wA","86LB_s1xsGA","ZtYbBOgMyhc","oM8pCGXKRzE","NdhWMqZEKJs","ASrAzDMdxXc","VRGzps8zD7A","PMInK_n7Mf8","saAir34l_LY","75SDRvl1CoQ","h1J7bPz2JNQ","b5imtlcwy-0","lLqg4W-SNOI","0xXAdDdzMuM","UXSFeosOQtE","hltaOKKshJM","Uz6nKoRNbfE","SyTk38rmggU","uUwDa1ZPcr0","QkrJvSKWaMs","Y2IizClYMSc","BjAvnakce2I","bv71OAbP5Os","PTNr5J5tDf0","fkGCtlWoOTc","MWkhH2xlT2E","BhO2nLDujUo","jDfx2tonErM","Y-tNj3uVwGw","aiU1hvJ4wyc","D6B2SR0CUXo","fTI_msRzTq4","Lc8ttSkn6bQ","B8FaKPLZwyI","6BXJUmnh3yw","Qzk3fZPpf_I","E8sFsrsjO6o","dj0LLXiIYGo","lcCMSiVFtPM","gCI4ceN4lE4","z0kJky4AU0Q","x54iuqJSUQA","tmxZnq2OKYI","TCWhf-ahLjA","CbJedBLifC4","0XfobrmSG0o","R65F9BZ4lkU","qC_mtrYLaQ8","hu-e3k2Fh6s","myPjHuAhhbE","4Zp-tjI6ucU","ZyKaHOnkVJM","K3Vs94y5rKI","Kmml3NkUMKo","-3J8ri9Uvys","UJZseul-Krw","t4R7RM76fQU","Wfcaus070N8","kSdGomHnxyQ","5Mxxm7SQOBE","ghIHtmmYwdI","UidF3dE-lgU","T6uwbmFi_yk","t3gI2S-IaeE","5B0U3JGRil0","IyrnMwcXw98","WqbETNWEnr0","woPBwUuRq3Q","e4jhUwZpJsA","-TUk8bCxC3M","bj2_GR0WPP0","40jxosIWAS8","tpSDtaNTtxw","44NBkkroYHY","0EpNumZ4I7g","eoOwxmINIRo","l3RGG-EXyz8","3lsZdWRF3RI","4U7qxEFf-D4","FYqUAbdRGEQ","wo0tWBb1JlM","eL33yN0xZEA","EtFDb4TWTdU","ogXsFR7gfis","0ti-88miKeI","OGfSjse2SBU","2J7jqDFfkFw","VKKT1iTmcG4","_jPTbJn48tU","hQreFlrxym4","G8HelpbzunI","f3xp29_hU-w","fHD8SJyg_v0","YftEb3BXt30","Bo8pEbYlBC0","31H1cYkeL48","iP35qaHicw8","sqnl3HkHVeE","zV6BY76juU4","ZhZH84hoNB4","mLCcxb0sBlc","9itOgrdraZM","C5ofu0K2Mrk","CD1EQY4EVNg","j368q0Jhxdc","jxlOOYz-xzk","D8tOEnZiGdQ","9ZANXY-Hk0Y","pNgNcrLqUpY","lvwDhAf9y_E","xG0mZi_vJeY","EuMr32jr2zU","hAC7SQFgpeA","h4RfdCpB4PI","uZcGgE1bSvk","ZL1Ig--LJ2g","twJqIuGx7RM","UcYkapy1q4w","_k3CLBaaxnE","7ldDv04lGOs","ePqQZ-pjGPs","tBv-TrFS6tE","z-ohqJnAkKo","bCBWDYTcAJQ","ssrOTE5G9jQ","JWlPH1LSAKQ","I1Tk_R5WpYg","gdXrpPqdUng","kPoqa7aTUiM","5LU5SmQs1WA","1zb7IzXfGgo","SslmB3D-ZzM","_d_7hzjOAck","j8uBkk4F7A4","Dsgt4YWxU8M","hUUhKkwQfSQ","tHql0-0mEjU","wYw1HyU8V3w","-kr8W7Xr4Fo","gOuNDKokKDA","D14MvGTEncM","VDiviS1tTW0","GVXp13lbj3A","lDH1z618OKo","VR6v4VcCDK8","8U1MJSvUzc0","-wUK5dPWhPk","rKaRniTUDkY","U5p2dMBk9Cw","UewHeYf27zk","-arIKVGYLZI","irNP2xXeRhk","PL4KBTPqdKk","nBJ3gfpRZ34","UT-_EXXH48w","ts_uEu595sY","LWAebG75lLU","1fnsToTH0S0","F5pYV0tyGW0","3zgzCIJHPJ4","Yr0KUculTZA","NO3VTXqMCoo","7ntIKlARh38","1y94BvawPwQ","4m8lZoyb6Xo","DbrC9axMUkg","RF_I8B_ap8o","2oVaQD8JXhA","xHxUerhIPAo","ylisULQJlJw","RKbDmeAEO2o","3ollHKfSKKw","JTDU-eKFpYU","p-bOdtJTnd8","KrxfYDpiN6Y","PFzGFOuAbxQ","qFfJGBpMj6A","-QYEP6uXlgw","uIAhE_8noaQ","SFrclMp4nng","0hyHNy7O_2s","-XwOTOwnm50","g2Heq1oStv8","6nDFRC8KmF8","rMxjDPOyjeY","bwv175AugYs","I8sxTCmirW0","hdGxz3E2qQg","plqYhI8tV2o","AiWk_pxkt4c","Eyvvq5enh8g","-s35Ja8wBwE","e9EqImvT8sA","d97V2DVJWCo","gxdzjnsHVDE","IkQRBhMeUhc","1RGPUJhPY4E","9jocO7MJMsY","UJmBUjFef1g","zNW3oz_j78Q","j56ZQOAuZJo","rtIs6bbpk2M","D5V2zHh8Das","jcV9R_Xm-_M","fVu-FPGW1ns","BHfccjHzKeo","cRAYmWv4-ME","Sm7SqRU6-Ec","yafwUp6FXQs","YuVGaK5VgZw","9BRP6FJ11OQ","jD_EXwcwiFk","_bkYgFyvwI8","Y7Q4U4Rtl-k","trUmUjsQb1k","FTW9CR9Y5bQ","a3CQB0Adghk","Z4oBzud283U","xMwFi72Q8mI","LQeXFsqLG-U","4VPiNdS5ntU","GRa1mzveZWM","snqighDpKjk","o8jniLW-oAA","BDl3vr9iSwo","PfpqId3Zc10","ZDPf_n-6MQI","UwEwcnuLhrQ","VQo2fa1eD_w","WaRU44d6POw","P4VlXa6l3Go","vDy-k6HUzzE","IC1lmLjVaOQ","4qQeU3IrEJY","vaDOshoGVXc","hkI85LHoQIA","Ag_GNms0Bfg","rZCeQIERGus","0dXRsh3uOoU","OYGcQW0hPYA","vN3SBokZAPw","_-HWHLObocg","mbGJ1iSuK_g","gvMqnqS4z_4","0SS0o9pYfKU","5Z42MvHUTME","t70l-9n1rCQ","-q2JXirzA8k","bKKQraCnui8","bfLyK_msMTY","SbUINENYUuQ","jIRSHLUK4Q4","nAxrW7axXxA","fKCjdcoin3A","7k7MQncsEug","zhjPd9or-oI","Da-XbQeaCeI","3QiRV4Hdm_g","GR3c1rVK0zY","25NJbpD60z8","kvDrCJcZGKc","Av8A_aqadd8","UIbMGs9pKm0","coL4ggcK6eE","rilHCNoV8vU","VgtQU1v3qj4","xdpJhCPCbz8","CNcKqb810jc","dBe8ZqVYkVM","1pvbMOPyLfo","Jjb5Oqoal3U","Din01zOTQJI","qu5FZa2Hgjc","RUXnNVf8IAw","o9auwgnBWz8","8-V8CnrEDgM","nn5XVHODK_A","bIjo6oZXm9c","jWUxyezueqc","Eedz7PHRvx0","rQcbWDe-y9s","i9qXdelj1B4","CtSdGrYrTp0","aJKSXQ48y-0","KcP3yp6w7Pw","cz-bUv8AqxA","TvYdeAnITLA","vnMXolJq1SA","J3kLwfj6DW4","0JYvG8E2yVM","q3u32Jp9mJ8","SCnXRBA7NC0","pzY4MCBZ8o0","CYRApZf_xdA","mwOv3B0gQo4","ZIFjYJYldVM","krHm441Bg5Y","zhyqkSj2Iag","n7kdKTHC5uQ","949lp2OVyY8","Pez0in248jg","Do5lpiArRBM","GNV7Teg-wIU","dOvTCcWpcR4","z553kEFrusI","X_YwB-1myGg","-SReSBGqjiQ","9HMlnppONyc","sW2gKAQL7-w","Pu72K46j69A","sEl8PFnlKWk","-y7Xi_kgryY","IWC369w5PqU","p-bMlegyUBU","ItUYKDE4o7M","rbzYOxXewaU","ILmzFpcETIs","ZPlkZDSaw6c","msAm_aNXiyY","rtZ6DSJ2faY","INQWH7mOGeQ","ZNFx5X0VEr0","8XIbgiNR6wo","37hS9KdZMfU","mlUGUZ0qTEE","47TMFC-iCCs","9cWlTf3i8hw","webTlt1a338","A2CsR3qeXg4","ojJpYgxknuQ","iOsp7lAhPtM","6Lls7rK2LsA","pW0rN5-oRmk","ArnYKNnSQaI","hXUUca1KgeE","cWtXE0zHxZg","VoB543m9Zzw","ZERWaMyGdF0","i7GCjzysvIw","cNNHh4jGQZE","3eX_JpmXONo","ljQb5b8-ME4","ibOta3IbNAY","HuQ-BfOHjtc","GgIWz18prRI","aA4IY3g1akY","sS3ttzIMkoI","OVEEsBW9P48","pnXXx00BxCk","GJgFtCk_H-g","pnyZVNcPoLE","XLnWFgX1-2Q","_MIMm3OCZXA","yTSL2rDHc6Y","U83GzcSCXRc","DmRIzMsJiLM","DUDs7F5Hhjo","8waq0d-ibCc","P2JUKJABpUM","TQkyG6EfIS4","M-S-dgxFs5I","1FnPqA1UqNA","vnb0wx-TF7s","bZEosed3jFA","DH8bqNsIw30","sOwcmcpAAkQ","LNizrobWjBU"]		
			
      var tag = document.createElement('script');

      tag.src = "https://www.youtube.com/iframe_api";
      var firstScriptTag = document.getElementsByTagName('script')[0];
      firstScriptTag.parentNode.insertBefore(tag, firstScriptTag);

      // 3. This function creates an <iframe> (and YouTube player)
      //    after the API code downloads.
      var player;
      function onYouTubeIframeAPIReady() {
        player = new YT.Player('youtubeContainer', {
          height: window.innerHeight - 25,
          width: window.innerWidth - 25,
		  videoId: urls[getRandomInt(0,urls.length - 1)],
          events: {
            'onReady': onPlayerReady,
            'onStateChange': onPlayerStateChange
          }
        });
      }
	  
      function onPlayerReady(event) {
        event.target.playVideo();
      }
	  
	  function onPlayerStateChange(event) {
	   if (event.data == YT.PlayerState.ENDED) {
		event.target.loadVideoById(urls[getRandomInt(0,urls.length - 1)]);
		}
      }
	  
	  function getRandomInt(min, max) {
		return Math.floor(Math.random() * (max - min)) + min;
		}
	  function itResizedOhShitKneeGrow() {
		player.setSize(window.innerWidth - 25, window.innerHeight - 25)
	  }
</script>
</body></html>`

func requestHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, html)
}

func main() {
	defer func() {
		fmt.Println(recover())
	}()
	/*dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	rawhtml, err := ioutil.ReadFile(strings.Join([]string{dir, "gamegramps.html"}, "/"))
	if err != nil {
		panic(err.Error())
	}
	html = string(rawhtml[:])*/
	fmt.Print("Hosting @ localhost:8080 gogo game groomples <3")
	http.HandleFunc("/gamegramps", requestHandler)
	http.ListenAndServe(":8080", nil)
}
