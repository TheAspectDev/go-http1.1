<?php

$hello = 5;


abstract class Fruit {
    public string $name;
    public string $description;
    public string $color;
    public string $icon;
    
    public function __construct() {
        $this->name = "";
        $this->description = "";
        $this->icon = "";
        $this->color = "";
    }
}
interface Juicy {
    public function getName(): string;
}
trait VeryJuicy {
    public function Squeeze () {
        echo "wuuoush " . $this->name;
    }
}
class Apple extends Fruit implements Juicy{
    use VeryJuicy;
    public function __construct() {
        $this->name = "Apple";
        $this->description = "Juicy";
        $this->icon = "A";
        $this->color = "red";
    }
    public function getName(): string {
        return $this->name;
    }
}
$fruit = new Apple();
echo $fruit->getName();
echo "\n
";
$fruit->Squeeze();

?>

hello