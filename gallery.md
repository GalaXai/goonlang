# example of filter use

<table>
<tr>
  <th>Original</th>
  <th>Downsampled</th>
  <th>Desaturated</th>
</tr>
<tr>
  <td><img src="static/galax.png" style="width: 200px; height: 200px" alt="Original"></td>
  <td><img src="static/downsapled.png" style="width: 200px; height: 200px" alt="Downsampled"></td>
  <td><img src="static/desaturated.png" style="width: 200px; height: 200px" alt="Desaturated"></td>
</tr>
<tr>
  <th>diffOfGaussians</th>
  <th>sobelFilter</th>
  <th>sobelFilter Gradients</th>
</tr>
<tr>
  <td><img src="static/gaussDiff.png" style="width: 200px; height: 200px" alt="difference of gaussians"></td>
  <td><img src="static/sobelFilter.png" style="width: 200px; height: 200px" alt="sobel filter"></td>
  <td><img src="static/gradientMatrix.png" style="width: 200px; height: 200px" alt="sobelFilter Gradients"></td>
</tr>
</table>

# ascii
> also inside bash code for colors
```bash
??????????????????????????????OOOOOOOO??????????????????????????
?????????????????????????????OPPPPoPPPOO????????????????????????
???????????????????????????OoccooPPPPoPOPP??????????????????????
????????????????????????PPocccccoooPPooooPOO????????????????????
???????????????OPocc;;ooooocccccccccoPPPPPOO????????????????????
??????????????occcc;,;cccoocc;ccccccoooPPPPOO???????????????????
????????????Occcc;;,;ccocccccccccccccccccoPPOOOOO???????????????
???????????Occcc;;,,ccccccccccccccccccccccPPOPOOO???????????????
????????PcOccccc;,;;ccccc;cccc;ccccccccccccoPPOOOOOOOO??????????
????????OPoccc;;,;;,ccccc;ccccccccccccccccccoPOOPPOO?PP?????????
?????????Occcc;;,;;;cc;cc;;ccccccccccccccccccPPOOPPOO?OP????????
?????????o;;;;;,;c;;cc;;;;;cccccccccccccccccccoPPPPPOOO?O???????
????????O;;c;;;;;;;;cc;;cc;c;cccccccccccccccccoooPoPPOO?O???????
????????c;;;;;,;cc;;;c;;;cc;;ccccccccccccccccccPPPPocO???O??????
???????c;;;;;;;;ccc;;c;;;ccc;;cccccccccccccccccoPPPPccOO?O??????
??????O;;;c;;;;;cccccc;;cccc;;;ccccccccccccccccoPPPPccoO?O??????
??????c;;;;;;;;;c;;ccccccccc;;;;ccccccccccccccccPOoP;;;P?O??????
?????c;;;;;;;;;;;c;cccccccccccc;;cccccccccccccccPPPoo;cc?O??????
????oc;;;c;;;;;;;;;ccccccccccccccccccccccccc;cccoooco;;cOO??????
???Occ;;;;;;c;;;;;;ccccc;cccccccc;ccccc;;cccc;cccocc;;;co???????
???occc;;;;;;;;;;;;ccc;;cccccccccc;ccccc;cc;;;cccc;;;;;cP???????
???cccc;;;;;;;;;;;;;cc;ccccccccccccccccc;;c;;;cc;c;;;;;cP???????
??Occcc;;;;;;;;;;;;;cc;cccccccccccccccccc;;;;;cc;o;;;;;cP???????
??Occc;;;;;;;;;;;;;;;c;cccccccccccocccccc;;ccccccc;,,,;cP???????
???ccc;;;c;;;;;;;;;;;;;cccccccccccccccccoccoccoccc;,,;;co???????
???oc;;;;c;;;;;;;;;;;;;cccccccccccccccccoccocccco;;,;,;cP???????
???Pc;;;;c;;;;;;;c;;;;;cccccccccccccccccoooPcoc;o,;;,,;cO???????
???O;;;;;;;;;,,;;;;;;;;;;;ccccccccccccoccoooooccc,;;;,;cO???????
????o;;;;;;;;,,;;;;;cccccccccccccoococoocoooOoc;,;,;;,;POO??????
?????cc;,,;;;;;;;c;;;cccccccccccooooooco;ocoOo;,,,;;;,;POOO?????
?????Oc;,,,;;;;;;;;;c;;;;cccccccoooPoooo;oPoPc;,,;;;;;cOOOO?????
??????P;,;,;c;;;;;;;;cc;;;cccccoccoPPPoooPPPPo;,;;;;;,cOOOO?????
???????P;;;,;;;;;;;;;;cccccooooocooPPOoPPPPOP;,,;;;;,;POOOO?????
????????P;;;,,,,;;;;;cc;ccoooooocooPOOPPPPOOO;;,,;;;;;OOOOO?????
?????????Oc;;;;;,;;;;;;ccooooooooooPOOPPPOOOP;;,;;;;;oOOOOOO????
???????????Pc;cccccc;;cooooooooPPPPPOOOPOOOOo;;,;;cc;OOOOOO?????
???????????oocccccc;;;PPPPPPPPPPPPPPOOPOOOOOPo;,;;;;oOOOOOO?????
???????????OPccoco;;cPOOPOOOOPOOOOOPOOPOOOOOOoc;;,;cPOOOOO??????
??????????????POOOPPPOO?OO???OOOOOOOOOOOOOOOOPPc,,;POOOO????????
??????????????POOOOOO?????????OOOOOOOOOOOOOOOOOo;;cOOO??????????
??????????????POOOOOO??????????OOOOOOOOOOOOOOOOOPcoOO???????????
??????????????POOOOOO???????????OOOOOO?OO?OOOOOOOPP?????????????
?????????????OOOOOOO??????????????OO???????OOOOOOPo?????????????
?????????????OOOOOOO??????????????????????????OOPPPO????????????
?????????????OOOOOOO??????????????????????????OOo;;;????????????
??????????????OOPPO??????????????????????????Oo;,,,,P???????????
?????????????????O???????????????????????????c,,,,,,;???????????
?????????????????O???????????????????????P;;;,,,,,,,,OO?????????
?????????????????OOOO???????????????????P,,,,,,,,,,,,oO?????????
???????????????????OPOO???????????????Oc,,,,,,,,,,,,,;??????????
???????????????????O?????????????OO??o;,,,,,,,,,,,,,,,P?????????
????????????????????OO??????????OOOP;,,,,,,,,,,,,,,,,,;O????????
???????????????????????????????OO??c,,,,,,,,,,,,,,,,,,,;????????
??????????????????????O???????OO???P,,,,,,,,,,,,,,,,,,,,o???????
??????????????????????O????????????O,,,,,,,,,,,,,,,,,,,,,P??????
??????????????????????????O????????Oo,,,,,,,,,,,,,,,,,,,,,P?????
???????????????????????????????????Oo,,,,,,,,,,,,,,,,,,,,,,cO???
???????????????????????????????????Oc,,,,,,,,,,,,,,,,,,,,,,,,cO?
???????????????????????????????????O,,,,,,,,,,,,,,,,,,,,,,,,,,,P
??????????????????????????????????O;,,,,,,,,,,,,,,,,,,,,,,,;;;;;
?????????????????????????????????P;,,,,,,,,,,,,,,,,,,,,,,,;;;,,;
???????????????????????????????O;,,,,,,,,,,,,,,,,,,,,,,,;;;;;;;,
??????????????????????????????P,,,,,,,,,,,,,,,,,,,,,,,;;;c;;;;;c
????????????????????????????Oc,,,,,,,,,,,,,,,,,,,,,,;;;c;,,;;,,,
```
## Cooking
```bash
                                                 \
                                                 \
                                                 \
                                                 \
                                                 \
                                                 \
                                                 \
                   \\\\\\\\\\\\                  \
                 \\\          \\\                \
               \\                \\              \
              \\                  \\\            \
            \\                      \\           \
           \\                        \\          \
           \                          \\         \
          \                            \         \
         \\                            \\        \
         \                              \        \
        \\                               \       \
        \                                \       \
        \                                \\      \
       \                                  \      \
       \                                  \      \
       \                                  \      \
       \                                  \      \
       \                                  \      \
       \                                  \      \
       \                                  \      \
       \                                  \      \
       \                                  \      \
       \\                                 \      \
        \                                \\      \
        \                                \       \
        \\                              \\       \
         \                              \        \
         \\                            \\        \
          \\                          \\         \
           \\                         \          \
            \\                       \           \
             \\                    \\            \
              \\                  \\             \
               \\\              \\\              \
                 \\\\        \\\\                \
                    \\\\\\\\\\                   \
                                                 \
                                                 \
                                                 \
                                                 \
                                                 \
                                                 \
```
