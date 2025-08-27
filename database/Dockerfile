FROM mysql:8.0

# Definir variáveis de ambiente
ENV MYSQL_ROOT_PASSWORD=minhasenha
ENV MYSQL_DATABASE=microservices

# Copiar scripts de inicialização
COPY init/ /docker-entrypoint-initdb.d/

# Configurações de performance
COPY my.cnf /etc/mysql/conf.d/

# Copiar script de health check
COPY healthcheck.sh /usr/local/bin/
RUN chmod +x /usr/local/bin/healthcheck.sh

# Criar diretório para logs
RUN mkdir -p /var/log/mysql && chown mysql:mysql /var/log/mysql

# Expor porta
EXPOSE 3306

# Healthcheck customizado
HEALTHCHECK --interval=30s --timeout=10s --start-period=60s --retries=3 \
  CMD /usr/local/bin/healthcheck.sh
