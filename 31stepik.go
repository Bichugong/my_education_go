package main

type MyStruct struct {
    On    bool
    Ammo  int
    Power int
}

func (m *MyStruct) Shoot() bool {
    if !m.On {
        return false
    }
    if m.Ammo > 0 {
        m.Ammo--
        return true
    }
    return false
}

func (m *MyStruct) RideBike() bool {
    if !m.On {
        return false
    }
    if m.Power > 0 {
        m.Power--
        return true
    }
    return false
}

func main() {
  testStruct := &MyStruct{
    On:    true,
    Ammo:  50,
    Power: 9,
   }
}
