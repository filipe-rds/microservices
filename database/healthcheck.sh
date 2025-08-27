#!/bin/bash
# =================================
# HEALTH CHECK PERSONALIZADO PARA MySQL
# =================================

# Verificar se o MySQL está respondendo
mysqladmin ping -h localhost -u root -p${MYSQL_ROOT_PASSWORD} --silent

if [ $? -eq 0 ]; then
    # Verificar se as databases principais existem
    mysql -u root -p${MYSQL_ROOT_PASSWORD} -e "USE \`order\`; USE \`payment\`; USE \`shipping\`;" 2>/dev/null
    
    if [ $? -eq 0 ]; then
        echo "MySQL healthy - all databases accessible"
        exit 0
    else
        echo "MySQL databases not ready"
        exit 1
    fi
else
    echo "MySQL not responding"
    exit 1
fi
