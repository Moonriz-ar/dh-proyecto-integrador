import styles from "./button.module.css";
import { clsx } from "clsx";

type Props = {
  variant?: "outlined" | "contained" | "text";
  children?: React.ReactNode;
  onClick: () => void;
};

function Button({ children, variant, onClick }: Props) {
  return (
    <button
      onClick={onClick}
      className={clsx({
        [styles.button]: true,
        [styles.contained]: variant === "contained" || variant === undefined,
        [styles.outlined]: variant === "outlined",
        [styles.text]: variant === "text",
      })}
    >
      {children}
    </button>
  );
}

export default Button;
