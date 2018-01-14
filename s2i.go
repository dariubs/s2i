package s2i

import(
  "strconv"
)

func Parse(strnum string, defaultnum interface{}) int64 {
  num, err := strconv.ParseInt(strnum, 10, 64)
  if err != nil {
    return int64(defaultnum.(int))
  }

  return num
}

func ParseUint(strnum string, defaultnum interface{}) uint {
  return uint(Parse(strnum, defaultnum))
}

func ParseInt(strnum string, defaultnum interface{}) int {
  return int(Parse(strnum, defaultnum))
}
