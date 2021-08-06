# 生成一个文件
pname="$1"
fname="$2"

dirInit (){
  dName1=$1
  dName2=$2
  if [ ! -d "${dName1}/${dName2}}" ]; then
    mkdir -p $dName1/$dName2

  fi
}

dirInit apis $pname
dirInit controllers $pname

