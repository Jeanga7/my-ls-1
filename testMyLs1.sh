#!/bin/bash

# Fonction pour tester et comparer les sorties
function test_ls {
    cmd_args="$@"
    echo "Testing: go run . $cmd_args"
    echo "-----------------------------------"
    
    # Exécute go run . et ls avec les mêmes arguments
    go run . $cmd_args > my_ls_output.txt
    ls $cmd_args > ls_output.txt
    
    # Compare les résultats
    if diff -q my_ls_output.txt ls_output.txt > /dev/null; then
        echo "Test passed!"
    else
        echo "Test failed. Output differs:"
        diff my_ls_output.txt ls_output.txt
    fi
    echo
    # Pause avant le prochain test, appuyer sur "Enter" pour continuer
    read -p "Press Enter to continue to the next test..."
}

# Nettoyage des fichiers de sortie de test
function cleanup {
    rm -f my_ls_output.txt ls_output.txt
}

# Tests sans arguments
test_ls

# Tests avec un fichier
file_name="testfile.txt"
touch $file_name
test_ls $file_name

# Tests avec un répertoire
dir_name="testdir"
mkdir $dir_name
test_ls $dir_name

# Tests avec le flag -l
test_ls -l
test_ls -l $file_name
test_ls -l $dir_name

# Tests avec le répertoire /usr/bin et les liens symboliques
test_ls -l /usr/bin

# Test du flag -R dans un répertoire contenant des sous-dossiers
mkdir -p testdir/subdir
touch testdir/file1.txt testdir/subdir/file2.txt
test_ls -R testdir

# Test du flag -a
test_ls -a

# Test du flag -r
test_ls -r

# Test du flag -t
test_ls -t

# Test du flag -la
test_ls -la

# Test avec les arguments "-l -t <directory>"
test_ls -l -t testdir

# Test avec les arguments "-lRr <directory>"
test_ls -lRr testdir

# Test avec les arguments "-l <directory> -a <file>"
test_ls -l $dir_name -a $file_name

# Test avec "-lR <directory>/<subdirectory>/"
test_ls -lR testdir/subdir///

# Test avec "-la /dev"
test_ls -la /dev

# Test avec "-alRrt <directory>"
test_ls -alRrt testdir

# Création d'un répertoire avec le nom "-"
mkdir ./-
test_ls -

# Création d'un fichier avec un lien symbolique
ln -s $file_name symlink_file
test_ls -l symlink_file/
test_ls -l symlink_file

# Création d'un répertoire avec des fichiers et un lien symbolique
ln -s $dir_name symlink_dir
test_ls -l symlink_dir/
test_ls -l symlink_dir

# Test de la performance avec la commande time
echo "Testing time performance with -R ~"
time go run . -R ~

# Nettoyage
cleanup

# Bonnes pratiques et tests de fichiers
echo "Checking code quality and tests..."
if [ -f "tests/test_my_ls.go" ]; then
    echo "Test file exists"
else
    echo "Test file is missing"
fi

echo "All tests completed."
