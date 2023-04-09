type Props = {
  firstName: string;
  lastName: string;
};

function User({ firstName, lastName }: Props) {
  const initials = firstName.slice(0, 1) + lastName.slice(0, 1);

  return (
    <section className="flex flex-col items-end md:flex-row md:gap-2 md:items-center">
      <div className="flex items-center justify-center w-10 h-10 mb-2 font-bold rounded-full bg-ivory text-secondary md:mb-0">
        <p>{initials}</p>
      </div>
      <div className="relative flex flex-col items-end md:items-start">
        <p className="text-sm">Hola, </p>
        <p className="text-sm font-bold text-secondary md:text-primary">{`${firstName} ${lastName}`}</p>
        <p className="absolute right-0 hidden font-bold -top-2 w-min h-min md:block">
          x
        </p>
      </div>
    </section>
  );
}

export default User;
