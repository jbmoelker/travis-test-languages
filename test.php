<?php

require_once 'vendor/autoload.php';

$loader = new Twig_Loader_Array(array(
    'index' => 'Hello from a {{ lang }} script!',
));
$twig = new Twig_Environment($loader);

echo $twig->render('index', array('lang' => 'PHP'));