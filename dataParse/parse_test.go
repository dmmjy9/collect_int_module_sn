package dataParse

import (
	"testing"
)

func TestParse(t *testing.T) {
	outStr := `
mujunyao@BJ-DHM-TC-ROUTER-01.BJ> show chassis hardware | no-more 
Hardware inventory:
Item             Version  Part number  Serial number     Description
Chassis                                JN1264925AFK      MX2010
Midplane         REV 58   750-044636   ABAC8765          Lower Backplane
Midplane 1       REV 02   711-044557   ABAC8553          Upper Backplane
PMP              REV 01   711-051406   ABAD2172          Power Midplane
FPM Board        REV 09   760-044634   ABDN8110          Front Panel Display
PSM 0            Rev 09   740-045051   1EDC621055S       AC 52V Power Supply Module
PSM 1            Rev 09   740-045051   1EDC62103RT       AC 52V Power Supply Module
PSM 2            Rev 09   740-045051   1EDC62103TX       AC 52V Power Supply Module
PSM 3            Rev 09   740-045051   1EDC621057M       AC 52V Power Supply Module
PSM 4            Rev 09   740-045051   1EDC621057D       AC 52V Power Supply Module
PSM 5            Rev 09   740-045051   1EDC62103S6       AC 52V Power Supply Module
PSM 6            Rev 09   740-045051   1EDC64206EZ       AC 52V Power Supply Module
PSM 7            Rev 09   740-045051   1EDC62103SB       AC 52V Power Supply Module
PSM 8            Rev 09   740-045051   1EDC62103TH       AC 52V Power Supply Module
PDM 0            REV 02   740-058680   1EFE6020059       AC SPH PDM
PDM 1            REV 02   740-058680   1EFE6020041       AC SPH PDM
Routing Engine 0 REV 06   740-041821   9016006679        RE-S-1800x4
Routing Engine 1 REV 06   740-041821   9016001905        RE-S-1800x4
CB 0             REV 08   750-063178   ABDN6845          Control Board
CB 1             REV 08   750-063178   ABDN6842          Control Board
SPMB 0           REV 05   711-041855   ABDN6782          PMB Board
SPMB 1           REV 05   711-041855   ABDK6997          PMB Board
SFB 0            REV 07   711-044466   ABDN7138          Switch Fabric Board
SFB 1            REV 07   711-044466   ABDN7176          Switch Fabric Board
SFB 2            REV 07   711-044466   ABDN7140          Switch Fabric Board
SFB 3            REV 07   711-044466   ABDN7184          Switch Fabric Board
SFB 4            REV 07   711-044466   ABDN7153          Switch Fabric Board
SFB 5            REV 07   711-044466   ABDJ2061          Switch Fabric Board
SFB 6            REV 07   711-044466   ABDN7118          Switch Fabric Board
SFB 7            REV 07   711-044466   ABDN7150          Switch Fabric Board
FPC 1            REV 72   750-044130   ABDN8487          MPC6E 3D
  CPU            REV 12   711-045719   ABDJ2544          RMPC PMB
  MIC 0          REV 19   750-049457   ABDN1680          2X100GE CFP2 OTN
    PIC 0                 BUILTIN      BUILTIN           2X100GE CFP2 OTN
      Xcvr 0     REV 01   740-065054   AF1543F004Y       CFP2-100G-SR10-D
      Xcvr 1              NON-JNPR     BP190403420003    100GBASE-LR4-D
  MIC 1          REV 28   750-046532   ABDD8448          24X10GE SFPP
    PIC 1                 BUILTIN      BUILTIN           24X10GE SFPP
      Xcvr 0     REV 01   740-031980   AA1619321AM       SFP+-10G-SR
      Xcvr 1     REV 01   740-031981   H4V2006918        SFP+-10G-LR
      Xcvr 2     REV 01   740-031980   AA1619321AD       SFP+-10G-SR
      Xcvr 3     REV 01   740-031980   AA1619321AF       SFP+-10G-SR
      Xcvr 4     REV 01   740-021309   G9A2006098        SFP+-10G-LR
      Xcvr 5     REV 01   740-031980   AA1619321GB       SFP+-10G-SR
      Xcvr 6     REV 01   740-031980   AA1619321GL       SFP+-10G-SR
      Xcvr 7     REV 01   740-031980   AA1619321GN       SFP+-10G-SR
      Xcvr 8     REV 01   740-031980   AA16193215H       SFP+-10G-SR
      Xcvr 9     REV 01   740-031980   AA16193215C       SFP+-10G-SR
      Xcvr 10    REV 01   740-031980   AA1619321H2       SFP+-10G-SR
      Xcvr 11    REV 01   740-031980   AA16193215A       SFP+-10G-SR
      Xcvr 12    REV 01   740-031980   AA16193215B       SFP+-10G-SR
      Xcvr 13    REV 01   740-031980   AD1710305PU       SFP+-10G-SR
      Xcvr 14    REV 01   740-031980   AA1619321GR       SFP+-10G-SR
      Xcvr 15    REV 01   740-031980   AA161932159       SFP+-10G-SR
      Xcvr 16             NON-JNPR     INFBT0062994      SFP+-10G-LR
      Xcvr 17    REV 01   740-031980   AA1619321B9       SFP+-10G-SR
      Xcvr 18    REV 01   740-021309   G3G2002917        SFP+-10G-LR
      Xcvr 19    REV 01   740-031980   AA161932158       SFP+-10G-SR
      Xcvr 20             NON-JNPR     MVK1VKK           SFP+-10G-SR
      Xcvr 21    REV 01   740-031980   AA1619321GG       SFP+-10G-SR
      Xcvr 22    REV 01   740-031980   AA1710304BS       SFP+-10G-SR
      Xcvr 23    REV 01   740-031980   AA1619321GK       SFP+-10G-SR
  XLM 0          REV 14   711-046638   ABDL0436          MPC6E XL
  XLM 1          REV 14   711-046638   ABDF8810          MPC6E XL
FPC 3            REV 72   750-044130   ABDN8362          MPC6E 3D
  CPU            REV 12   711-045719   ABDL8702          RMPC PMB
  MIC 0          REV 19   750-049457   ABDJ2341          2X100GE CFP2 OTN
    PIC 0                 BUILTIN      BUILTIN           2X100GE CFP2 OTN
      Xcvr 0     REV 01   740-065054   AF1543F004V       CFP2-100G-SR10-D
      Xcvr 1     REV 01   740-065054   AF1542F008C       CFP2-100G-SR10-D
  MIC 1          REV 33   750-046532   CAHV4129          24X10GE SFPP
    PIC 1                 BUILTIN      BUILTIN           24X10GE SFPP
      Xcvr 0     REV 01   740-031980   AD1710305S1       SFP+-10G-SR
      Xcvr 1     REV 01   740-031980   AA1703302LL       SFP+-10G-SR
      Xcvr 2     REV 01   740-031980   AD1710305G0       SFP+-10G-SR
      Xcvr 3     REV 01   740-021308   CG09KN0WK         SFP+-10G-SR
      Xcvr 4     REV 01   740-031980   AA1710304CP       SFP+-10G-SR
      Xcvr 5     REV 01   740-031980   AD1710305G2       SFP+-10G-SR
      Xcvr 6              NON-JNPR     T15L85393         SFP+-10G-ER
      Xcvr 7     REV 01   740-031980   AD1710305PD       SFP+-10G-SR
      Xcvr 8     REV 01   740-031980   AA1703301ZB       SFP+-10G-SR
      Xcvr 9     REV 01   740-031980   AA1703301ZG       SFP+-10G-SR
      Xcvr 10             NON-JNPR     AW909F0           SFP+-10G-LR
      Xcvr 11    REV 01   740-031980   AD1710305PM       SFP+-10G-SR
      Xcvr 12    REV 01   740-031980   AD1710305R5       SFP+-10G-SR
      Xcvr 13    REV 01   740-031980   AA1703301YX       SFP+-10G-SR
      Xcvr 14    REV 01   740-031980   AA1703302KP       SFP+-10G-SR
      Xcvr 15             NON-JNPR     T15H37173         SFP+-10G-ER
      Xcvr 16    REV 01   740-031980   AA1710304BR       SFP+-10G-SR
      Xcvr 17             NON-JNPR     HA18500531592     SFP+-10G-SR
      Xcvr 18    REV 01   740-031980   AA1710304BU       SFP+-10G-SR
      Xcvr 19    REV 01   740-031980   AA1703301Z9       SFP+-10G-SR
      Xcvr 20    REV 01   740-031980   AD1710305RH       SFP+-10G-SR
      Xcvr 21    REV 01   740-031980   AD1710305SB       SFP+-10G-SR
      Xcvr 22    REV 01   740-031980   AA1703302L0       SFP+-10G-SR
      Xcvr 23    REV 01   740-031980   AD1710305P5       SFP+-10G-SR
  XLM 0          REV 14   711-046638   ABDN6162          MPC6E XL
  XLM 1          REV 14   711-046638   ABDN6238          MPC6E XL
FPC 5            REV 56   750-054576   CAGG5191          MPC8E 3D
  CPU            REV 19   750-057177   CAGY2796          SMPC PMB
  MIC 0          REV 13   750-055992   CAHL3853          MRATE-12xQSFPP-XGE-XLGE-CGE
    PIC 0                 BUILTIN      BUILTIN           MRATE-12xQSFPP-XGE-XLGE-CGE
      Xcvr 0     REV 01   740-046565   QH11035N          QSFP+-40G-SR4
      Xcvr 1     REV 01   740-046565   QH1103B3          QSFP+-40G-SR4
      Xcvr 2     REV 01   740-046565   QH1103B0          QSFP+-40G-SR4
      Xcvr 3     REV 01   740-046565   QH110372          QSFP+-40G-SR4
      Xcvr 4     REV 01   740-046565   QH11035W          QSFP+-40G-SR4
      Xcvr 5     REV 01   740-046565   QH11037H          QSFP+-40G-SR4
      Xcvr 6              NON-JNPR     U2292H002HJ       QSFP-100GBASE-SR4
      Xcvr 7              NON-JNPR     U2292H002RP       QSFP-100GBASE-SR4
      Xcvr 8     REV 01   740-046565   QH11035R          QSFP+-40G-SR4
      Xcvr 9     REV 01   740-046565   QH11037D          QSFP+-40G-SR4
      Xcvr 10    REV 01   740-046565   QH11035V          QSFP+-40G-SR4
      Xcvr 11    REV 01   740-046565   QH11035J          QSFP+-40G-SR4
FPC 7            REV 48   750-053323   CALA2404          MPC7E 3D 40XGE
  CPU            REV 21   750-057177   CALB0003          SMPC PMB
  PIC 0                   BUILTIN      BUILTIN           20x10GE SFPP
    Xcvr 0       REV 01   740-021308   CG09KN0WD         SFP+-10G-SR
    Xcvr 1       REV 01   740-021308   CG09KN0WV         SFP+-10G-SR
    Xcvr 2       REV 01   740-021309   G9A2006082        SFP+-10G-LR
    Xcvr 3       REV 01   740-021308   CG09KN0QT         SFP+-10G-SR
    Xcvr 4       REV 01   740-021309   G9A2006080        SFP+-10G-LR
    Xcvr 5       REV 01   740-021308   CG09KN1PG         SFP+-10G-SR
    Xcvr 6       REV 01   740-021308   CG09KN0PX         SFP+-10G-SR
    Xcvr 7       REV 01   740-021308   CG09KN0PK         SFP+-10G-SR
    Xcvr 8       REV 01   740-030128   T10C81812         SFP+-10G-ER
    Xcvr 9       REV 01   740-021309   G9A2006084        SFP+-10G-LR
    Xcvr 10      REV 01   740-021309   G9A2006092        SFP+-10G-LR
    Xcvr 11      REV 01   740-021308   CG09KN0PW         SFP+-10G-SR
    Xcvr 12      REV 01   740-021309   G9A2006094        SFP+-10G-LR
    Xcvr 13      REV 01   740-031980   AA1619321H3       SFP+-10G-SR
    Xcvr 14      REV 01   740-021308   CG09KN0WY         SFP+-10G-SR
    Xcvr 15      REV 01   740-021308   CG09KN1PD         SFP+-10G-SR
    Xcvr 16      REV 01   740-021308   CG09KN1PE         SFP+-10G-SR
    Xcvr 17      REV 01   740-021308   CG09KN0WP         SFP+-10G-SR
    Xcvr 18      REV 01   740-021308   CG09KN0WB         SFP+-10G-SR
    Xcvr 19      REV 01   740-021308   CG09KN0PF         SFP+-10G-SR
  PIC 1                   BUILTIN      BUILTIN           20x10GE SFPP
    Xcvr 0       REV 01   740-021308   CG09KN0PB         SFP+-10G-SR
    Xcvr 1       REV 01   740-021308   CG09KN1PC         SFP+-10G-SR
    Xcvr 2       REV 01   740-021308   CG09KN0PP         SFP+-10G-SR
    Xcvr 3       REV 01   740-021308   CG09KN0PA         SFP+-10G-SR
    Xcvr 4       REV 01   740-021308   CG09KN0WN         SFP+-10G-SR
    Xcvr 5       REV 01   740-021308   CG09KN0PH         SFP+-10G-SR
    Xcvr 6       REV 01   740-021308   CG09KN0WX         SFP+-10G-SR
    Xcvr 7       REV 01   740-021309   G9A2006088        SFP+-10G-LR
    Xcvr 8       REV 01   740-021308   CG09KN0WR         SFP+-10G-SR
    Xcvr 9       REV 01   740-021309   G9A2006090        SFP+-10G-LR
    Xcvr 10      REV 01   740-021309   G9A2006086        SFP+-10G-LR
    Xcvr 11      REV 01   740-031980   AA1710304BN       SFP+-10G-SR
    Xcvr 12      REV 01   740-031981   UHQ027X           SFP+-10G-LR
    Xcvr 13      REV 01   740-021308   CG09KN0PR         SFP+-10G-SR
    Xcvr 14      REV 01   740-021309   G9A2006100        SFP+-10G-LR
    Xcvr 15      REV 01   740-021309   G9A2006096        SFP+-10G-LR
    Xcvr 16      REV 01   740-041612   Z14280HLR         SFP+-10G-ZR
    Xcvr 17      REV 01   740-030128   T10C81810         SFP+-10G-ER
    Xcvr 18      REV 01   740-041612   Z14280HLN         SFP+-10G-ZR
    Xcvr 19      REV 01   740-021308   CG09KN0WW         SFP+-10G-SR
  WAN MEZZ       REV 11   750-055620   CALE3508          MPC7E 40x10GE Mezz
FPC 9            REV 46   750-056519   CALL5432          MPC7E 3D MRATE-12xQSFPP-XGE-XLGE-CGE
  CPU            REV 21   750-057177   CALT3035          SMPC PMB
  PIC 0                   BUILTIN      BUILTIN           MRATE-6xQSFPP-XGE-XLGE-CGE
    Xcvr 2       REV 01   740-058734   1ECQ123702S       QSFP-100GBASE-SR4
    Xcvr 5       REV 01   740-058734   1ECQ1237030       QSFP-100GBASE-SR4
  PIC 1                   BUILTIN      BUILTIN           MRATE-6xQSFPP-XGE-XLGE-CGE
    Xcvr 2       REV 01   740-058734   1ECQ123701B       QSFP-100GBASE-SR4
    Xcvr 5       REV 01   740-058734   1ECQ123708G       QSFP-100GBASE-SR4
ADC 7            REV 17   750-043596   ABCD5405          Adapter Card
ADC 9            REV 23   750-043596   CAMG4030          Adapter Card
Fan Tray 0       REV 03   760-052467   ACAY8815          172mm FanTray - 6 Fans
Fan Tray 1       REV 03   760-052467   ACAY8802          172mm FanTray - 6 Fans
Fan Tray 2       REV 03   760-052467   ACAY8776          172mm FanTray - 6 Fans
Fan Tray 3       REV 03   760-052467   ACAY8770          172mm FanTray - 6 Fans

{master}
mujunyao@BJ-DHM-TC-ROUTER-01.BJ> exit 

Connection to 172.29.254.3 closed.

`
	_ = JuniperParser(outStr)
	//for _, val := range info {
	//	fmt.Printf("%d -- %s -- %s -- %s\n", val.Id, val.IntName, val.ModType, val.ModSn)
	//}
}
